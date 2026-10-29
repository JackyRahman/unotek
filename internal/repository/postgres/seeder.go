package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/internal/domain"
	"unotek-backend/pkg/security"
)

func SeedInitialData(db *gorm.DB) error {
	ctx := context.Background()

	// Fast-check: Jika data admin sudah ada di database, lewati proses seeding agar startup instan
	var adminCount int64
	if err := db.WithContext(ctx).Model(&domain.User{}).Where("email = ?", "admin@company.com").Count(&adminCount).Error; err == nil && adminCount > 0 {
		log.Println("[SEEDER] Database sudah ter-seed sebelumnya. Melewati seeding untuk startup cepat.")
		return nil
	}

	log.Println("[SEEDER] Checking and seeding initial master data...")

	// 1. Roles
	superAdminRoleID := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	adminRoleID := uuid.MustParse("22222222-2222-2222-2222-222222222222")
	editorRoleID := uuid.MustParse("33333333-3333-3333-3333-333333333333")

	roles := []domain.Role{
		{
			ID:          superAdminRoleID,
			Name:        "super_admin",
			DisplayName: "Super Administrator",
			Description: "Full system access and privileges",
		},
		{
			ID:          adminRoleID,
			Name:        "admin",
			DisplayName: "Administrator",
			Description: "Standard administrative operations",
		},
		{
			ID:          editorRoleID,
			Name:        "editor",
			DisplayName: "Content Editor",
			Description: "Manage content and master data",
		},
	}

	for _, role := range roles {
		var existing domain.Role
		if err := db.WithContext(ctx).Where("name = ?", role.Name).First(&existing).Error; err != nil {
			if err := db.WithContext(ctx).Create(&role).Error; err != nil {
				log.Printf("[SEEDER] Error creating role %s: %v\n", role.Name, err)
			}
		}
	}

	// 2. Permissions
	permissions := []domain.Permission{
		{Name: "users:create", Resource: "users", Action: "create", Description: "Create new users"},
		{Name: "users:read", Resource: "users", Action: "read", Description: "View user profiles"},
		{Name: "users:update", Resource: "users", Action: "update", Description: "Modify user details"},
		{Name: "users:delete", Resource: "users", Action: "delete", Description: "Remove users"},

		{Name: "roles:read", Resource: "roles", Action: "read", Description: "View roles and permissions"},
		{Name: "roles:write", Resource: "roles", Action: "write", Description: "Assign and update roles"},

		{Name: "clients:create", Resource: "clients", Action: "create", Description: "Add new clients"},
		{Name: "clients:read", Resource: "clients", Action: "read", Description: "View clients list and details"},
		{Name: "clients:update", Resource: "clients", Action: "update", Description: "Update client information"},
		{Name: "clients:delete", Resource: "clients", Action: "delete", Description: "Delete clients"},

		{Name: "partners:create", Resource: "partners", Action: "create", Description: "Add new partners"},
		{Name: "partners:read", Resource: "partners", Action: "read", Description: "View partners list and details"},
		{Name: "partners:update", Resource: "partners", Action: "update", Description: "Update partner information"},
		{Name: "partners:delete", Resource: "partners", Action: "delete", Description: "Delete partners"},

		{Name: "solutions:create", Resource: "solutions", Action: "create", Description: "Add solutions and features"},
		{Name: "solutions:read", Resource: "solutions", Action: "read", Description: "View solutions and features"},
		{Name: "solutions:update", Resource: "solutions", Action: "update", Description: "Update solutions"},
		{Name: "solutions:delete", Resource: "solutions", Action: "delete", Description: "Delete solutions"},

		{Name: "team:create", Resource: "team", Action: "create", Description: "Add team members"},
		{Name: "team:read", Resource: "team", Action: "read", Description: "View team members"},
		{Name: "team:update", Resource: "team", Action: "update", Description: "Update team members"},
		{Name: "team:delete", Resource: "team", Action: "delete", Description: "Delete team members"},

		{Name: "contacts:read", Resource: "contacts", Action: "read", Description: "View submitted contact messages"},
		{Name: "contacts:update", Resource: "contacts", Action: "update", Description: "Change contact message status"},
		{Name: "contacts:delete", Resource: "contacts", Action: "delete", Description: "Delete contact submissions"},

		{Name: "company:update", Resource: "company", Action: "update", Description: "Update company profile and offices"},
	}

	for _, perm := range permissions {
		var existing domain.Permission
		if err := db.WithContext(ctx).Where("name = ?", perm.Name).First(&existing).Error; err != nil {
			db.WithContext(ctx).Create(&perm)
		}
	}

	// 3. Assign all permissions to super_admin
	var allPermissions []domain.Permission
	db.WithContext(ctx).Find(&allPermissions)
	for _, p := range allPermissions {
		var count int64
		db.Table("role_permissions").Where("role_id = ? AND permission_id = ?", superAdminRoleID, p.ID).Count(&count)
		if count == 0 {
			db.Table("role_permissions").Create(&domain.RolePermission{
				RoleID:       superAdminRoleID,
				PermissionID: p.ID,
			})
		}
	}

	// 4. Seed Super Admin User (admin@company.com / admin12345)
	var existingUser domain.User
	if err := db.WithContext(ctx).Where("email = ?", "admin@company.com").First(&existingUser).Error; err != nil {
		passwordHash, err := security.HashPassword("admin12345")
		if err != nil {
			return err
		}

		superAdminUser := domain.User{
			ID:           uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Email:        "admin@company.com",
			Username:     "superadmin",
			PasswordHash: passwordHash,
			FullName:     "System Super Administrator",
			Phone:        "+628110000123",
			IsActive:     true,
		}

		if err := db.WithContext(ctx).Create(&superAdminUser).Error; err == nil {
			db.Table("user_roles").Create(&domain.UserRole{
				UserID: superAdminUser.ID,
				RoleID: superAdminRoleID,
			})
			log.Println("[SEEDER] Super Admin account created: admin@company.com / admin12345")
		}
	}

	// 5. Company Profile
	var profileCount int64
	db.WithContext(ctx).Model(&domain.CompanyProfile{}).Count(&profileCount)
	if profileCount == 0 {
		profile := domain.CompanyProfile{
			ID:              uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			CompanyName:     "Unotek Digital Solusindo",
			Tagline:         "Empowering Enterprise Logistics & Supply Chain with Next-Gen Software",
			AboutText:       "Kami adalah software house spesialis enterprise solutions yang berfokus pada efisiensi supply chain, smart warehousing, dan automasi proses bisnis terintegrasi.",
			Vision:          "Menjadi mitra transformasi digital nomor satu di Asia Tenggara untuk solusi industri & logistik terintegrasi.",
			Mission:         "Mengembangkan produk perangkat lunak berkualitas tinggi, handal, dan berkinerja tinggi yang memberikan dampak bisnis nyata bagi setiap mitra dan klien.",
			FoundedYear:     2018,
			LogoURL:         "https://assets.company.com/brand/unotek-logo.svg",
			FaviconURL:      "https://assets.company.com/brand/favicon.ico",
			MetaTitle:       "Unotek - Enterprise Logistics & Supply Chain Software Solutions",
			MetaDescription: "Partner software house terpercaya untuk Warehouse Management System (WMS), Transport Management System (TMS), dan automasi industri.",
		}
		db.WithContext(ctx).Create(&profile)
	}

	// 6. Office Locations
	var officeCount int64
	db.WithContext(ctx).Model(&domain.OfficeLocation{}).Count(&officeCount)
	if officeCount == 0 {
		offices := []domain.OfficeLocation{
			{
				BranchName:   "Jakarta Head Office",
				BranchType:   "head_office",
				Address:      "Menara Standard Chartered Lt. 28, Jl. Prof. DR. Satrio No.164",
				City:         "Jakarta Selatan",
				Province:     "DKI Jakarta",
				PostalCode:   "12930",
				Country:      "Indonesia",
				Phone:        "+62 21 5550 1928",
				Email:        "contact@company.com",
				WorkingHours: "Senin - Jumat: 08:30 - 17:30 WIB",
				IsHeadOffice: true,
				DisplayOrder: 1,
				IsActive:     true,
			},
			{
				BranchName:   "Bandung Tech Hub",
				BranchType:   "development_center",
				Address:      "Dago Digital Valley, Jl. Ir. H. Juanda No. 84",
				City:         "Bandung",
				Province:     "Jawa Barat",
				PostalCode:   "40132",
				Country:      "Indonesia",
				Phone:        "+62 22 7200 4811",
				Email:        "dev@company.com",
				WorkingHours: "Senin - Jumat: 09:00 - 18:00 WIB",
				IsHeadOffice: false,
				DisplayOrder: 2,
				IsActive:     true,
			},
		}
		for _, o := range offices {
			db.WithContext(ctx).Create(&o)
		}
	}

	// 7. Clients
	var clientCount int64
	db.WithContext(ctx).Model(&domain.Client{}).Count(&clientCount)
	if clientCount == 0 {
		clients := []domain.Client{
			{
				Name:         "PT Indomarco Prismatama (Indomaret)",
				Slug:         "indomaret",
				Industry:     "Retail & FMCG",
				WebsiteURL:   "https://indomaret.co.id",
				LogoURL:      "https://assets.company.com/clients/indomaret.svg",
				Description:  "Implementasi sistem monitoring distribusi toko dan inventori regional.",
				IsFeatured:   true,
				DisplayOrder: 1,
				IsActive:     true,
			},
			{
				Name:         "PT Sumber Alfaria Trijaya Tbk (Alfamart)",
				Slug:         "alfamart",
				Industry:     "Retail & Distribution",
				WebsiteURL:   "https://alfamart.co.id",
				LogoURL:      "https://assets.company.com/clients/alfamart.svg",
				Description:  "Sistem integrasi warehouse batching and sorting.",
				IsFeatured:   true,
				DisplayOrder: 2,
				IsActive:     true,
			},
			{
				Name:         "Pos Indonesia Logistics",
				Slug:         "pos-indonesia",
				Industry:     "Postal & Logistics",
				WebsiteURL:   "https://posindonesia.co.id",
				LogoURL:      "https://assets.company.com/clients/pos-indonesia.svg",
				Description:  "Modernisasi track & trace dan warehouse hub nasional.",
				IsFeatured:   true,
				DisplayOrder: 3,
				IsActive:     true,
			},
			{
				Name:         "DHL Supply Chain Indonesia",
				Slug:         "dhl-indonesia",
				Industry:     "Global Logistics",
				WebsiteURL:   "https://dhl.com",
				LogoURL:      "https://assets.company.com/clients/dhl.svg",
				Description:  "Automasi manajemen cross-docking dan multi-client warehouse.",
				IsFeatured:   true,
				DisplayOrder: 4,
				IsActive:     true,
			},
			{
				Name:         "Kalbe Farma Distribution",
				Slug:         "kalbe-farma",
				Industry:     "Pharmaceutical",
				WebsiteURL:   "https://kalbe.co.id",
				LogoURL:      "https://assets.company.com/clients/kalbe.svg",
				Description:  "Cold chain warehouse compliance dan barcode validation system.",
				IsFeatured:   true,
				DisplayOrder: 5,
				IsActive:     true,
			},
		}
		for _, c := range clients {
			db.WithContext(ctx).Create(&c)
		}
	}

	// 8. Product Partners
	var partnerCount int64
	db.WithContext(ctx).Model(&domain.ProductPartner{}).Count(&partnerCount)
	if partnerCount == 0 {
		partners := []domain.ProductPartner{
			{
				Name:         "Zebra Technologies",
				Slug:         "zebra-technologies",
				PartnerType:  "hardware",
				WebsiteURL:   "https://zebra.com",
				LogoURL:      "https://assets.company.com/partners/zebra.svg",
				Description:  "Penyedia barcode scanner, mobile computer enterprise, dan RFID hardware terkemuka.",
				DisplayOrder: 1,
				IsActive:     true,
			},
			{
				Name:         "Honeywell Safety & Productivity",
				Slug:         "honeywell",
				PartnerType:  "hardware",
				WebsiteURL:   "https://honeywell.com",
				LogoURL:      "https://assets.company.com/partners/honeywell.svg",
				Description:  "Partner resmi untuk supply chain mobility device dan industrial thermal printer.",
				DisplayOrder: 2,
				IsActive:     true,
			},
			{
				Name:         "Amazon Web Services (AWS)",
				Slug:         "amazon-web-services",
				PartnerType:  "cloud",
				WebsiteURL:   "https://aws.amazon.com",
				LogoURL:      "https://assets.company.com/partners/aws.svg",
				Description:  "Infrastruktur cloud enterprise dengan ketersediaan tinggi dan sertifikasi ISO 27001.",
				DisplayOrder: 3,
				IsActive:     true,
			},
			{
				Name:         "SAP Integration Partner",
				Slug:         "sap",
				PartnerType:  "enterprise_software",
				WebsiteURL:   "https://sap.com",
				LogoURL:      "https://assets.company.com/partners/sap.svg",
				Description:  "Konektivitas seamless dengan SAP S/4HANA & ECC untuk modul MM dan SD.",
				DisplayOrder: 4,
				IsActive:     true,
			},
			{
				Name:         "Hikvision Industrial Automation",
				Slug:         "hikvision",
				PartnerType:  "iot_automation",
				WebsiteURL:   "https://hikvision.com",
				LogoURL:      "https://assets.company.com/partners/hikvision.svg",
				Description:  "Partner solusi computer vision dan OCR gate system untuk pergudangan.",
				DisplayOrder: 5,
				IsActive:     true,
			},
		}
		for _, p := range partners {
			db.WithContext(ctx).Create(&p)
		}
	}

	// 9. Solutions with Features and Modules
	var solutionCount int64
	db.WithContext(ctx).Model(&domain.Solution{}).Count(&solutionCount)
	if solutionCount == 0 {
		wmsID := uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc")
		wms := domain.Solution{
			ID:           wmsID,
			Title:        "Warehouse Management System (WMS)",
			Slug:         "warehouse-management-system",
			Tagline:      "Smart, Scalable, and Real-Time Warehouse Control",
			Summary:      "Sistem otomasi gudang modern yang mengoptimalkan seluruh pergerakan barang mulai dari receiving, putaway, picking, packing hingga dispatch secara real-time.",
			Description:  "WMS kami dirancang khusus untuk menangani kompleksitas pergudangan high-throughput dengan akurasi stok hingga 99.9%. Dilengkapi integrasi barcode, barcode scanner 2D, mobile RF, dan algoritma putaway/picking otomatis untuk memangkas waktu operasional hingga 45%.",
			IconURL:      "https://assets.company.com/icons/wms-box.svg",
			HeroImageURL: "https://assets.company.com/mockups/wms-hero-dashboard.png",
			IsFeatured:   true,
			DisplayOrder: 1,
			IsActive:     true,
			Features: []domain.SolutionFeature{
				{
					Title:        "Real-Time Inventory Visibility",
					Description:  "Pantau kuantiti dan lokasi stok di berbagai gudang hingga level bin/rak dengan pembaruan instan.",
					IconURL:      "https://assets.company.com/icons/eye.svg",
					DisplayOrder: 1,
					IsActive:     true,
				},
				{
					Title:        "Intelligent Wave & Batch Picking",
					Description:  "Algoritma rute picking terpendek untuk meningkatkan kecepatan pemenuhan order secara drastis.",
					IconURL:      "https://assets.company.com/icons/route.svg",
					DisplayOrder: 2,
					IsActive:     true,
				},
				{
					Title:        "Multi-Warehouse & Multi-Client",
					Description:  "Kelola banyak entitas gudang dan 3PL multi-tenancy dalam satu dashboard terpusat.",
					IconURL:      "https://assets.company.com/icons/layers.svg",
					DisplayOrder: 3,
					IsActive:     true,
				},
				{
					Title:        "Mobile Handheld Scanner Ready",
					Description:  "Aplikasi mobile Android responsif untuk scanning barcode 1D/2D dan QR Code secara offline/online.",
					IconURL:      "https://assets.company.com/icons/scanner.svg",
					DisplayOrder: 4,
					IsActive:     true,
				},
			},
			Modules: []domain.SolutionModule{
				{
					Title:        "Inbound & Putaway Automation",
					Description:  "Penerimaan PO barang, validasi kuantiti, pencetakan barcode label otomatis, dan rekomendasi rak kosong.",
					ImageURL:     "https://assets.company.com/modules/inbound.png",
					DisplayOrder: 1,
					IsActive:     true,
				},
				{
					Title:        "Inventory & Cycle Count Control",
					Description:  "Stock opname berkala tanpa perlu menghentikan operasional gudang, tracking lot number, dan tanggal kadaluarsa (FIFO/FEFO).",
					ImageURL:     "https://assets.company.com/modules/inventory.png",
					DisplayOrder: 2,
					IsActive:     true,
				},
				{
					Title:        "Outbound & Dispatch Management",
					Description:  "Verifikasi packing sebelum muat truk, integrasi kurir, dan cetak Surat Jalan otomatis.",
					ImageURL:     "https://assets.company.com/modules/outbound.png",
					DisplayOrder: 3,
					IsActive:     true,
				},
			},
		}
		db.WithContext(ctx).Create(&wms)

		tmsID := uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd")
		tms := domain.Solution{
			ID:           tmsID,
			Title:        "Transport Management System (TMS)",
			Slug:         "transport-management-system",
			Tagline:      "End-to-End Fleet Optimization & Delivery Tracking",
			Summary:      "Solusi manajemen armada pengiriman, optimalisasi rute cerdas (route planner), dan proof-of-delivery (e-POD) digital berbasis mobile app driver.",
			Description:  "Mengintegrasikan perencanaan armada dengan telematika GPS untuk meminimalkan biaya bahan bakar, mengontrol lead time pengiriman, dan memberikan tracking publik real-time kepada pelanggan.",
			IconURL:      "https://assets.company.com/icons/tms-truck.svg",
			HeroImageURL: "https://assets.company.com/mockups/tms-hero-map.png",
			IsFeatured:   true,
			DisplayOrder: 2,
			IsActive:     true,
		}
		db.WithContext(ctx).Create(&tms)
	}

	// 10. Team Members
	var teamCount int64
	db.WithContext(ctx).Model(&domain.TeamMember{}).Count(&teamCount)
	if teamCount == 0 {
		members := []domain.TeamMember{
			{
				FullName:     "Arya Wibowo",
				Position:     "Chief Executive Officer",
				Department:   "Executive",
				Bio:          "15+ tahun pengalaman memimpin transformasi sistem digital supply chain di berbagai perusahaan multinasional.",
				PhotoURL:     "https://assets.company.com/team/arya.jpg",
				LinkedinURL:  "https://linkedin.com/in/aryawibowo",
				DisplayOrder: 1,
				IsLeadership: true,
				IsActive:     true,
			},
			{
				FullName:     "Siti Rahmawati",
				Position:     "Chief Technology Officer",
				Department:   "Technology",
				Bio:          "Mantan Principal Software Architect dengan spesialisasi distributed systems, high throughput event-driven architecture, dan cloud security.",
				PhotoURL:     "https://assets.company.com/team/siti.jpg",
				LinkedinURL:  "https://linkedin.com/in/sitirahmawati",
				DisplayOrder: 2,
				IsLeadership: true,
				IsActive:     true,
			},
			{
				FullName:     "Budi Santoso",
				Position:     "Head of Product Development",
				Department:   "Product",
				Bio:          "Berpengalaman mendesain produk enterprise B2B dan solusi IoT supply chain berorientasi efisiensi operasional.",
				PhotoURL:     "https://assets.company.com/team/budi.jpg",
				LinkedinURL:  "https://linkedin.com/in/budisantoso",
				DisplayOrder: 3,
				IsLeadership: true,
				IsActive:     true,
			},
			{
				FullName:     "Dimas Pratama",
				Position:     "Lead Solutions Architect",
				Department:   "Engineering",
				Bio:          "Spesialis integrasi ERP (SAP, Oracle, NetSuite) dan pengembangan core engine pergudangan modern.",
				PhotoURL:     "https://assets.company.com/team/dimas.jpg",
				LinkedinURL:  "https://linkedin.com/in/dimaspratama",
				DisplayOrder: 4,
				IsLeadership: false,
				IsActive:     true,
			},
		}
		for _, m := range members {
			db.WithContext(ctx).Create(&m)
		}
	}

	log.Println("[SEEDER] Initial data seeding finished successfully.")
	return nil
}
