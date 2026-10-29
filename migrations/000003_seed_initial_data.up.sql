-- 1. Seed Roles
INSERT INTO roles (id, name, display_name, description) VALUES
('11111111-1111-1111-1111-111111111111', 'super_admin', 'Super Administrator', 'Full system access and privileges'),
('22222222-2222-2222-2222-222222222222', 'admin', 'Administrator', 'Standard administrative operations'),
('33333333-3333-3333-3333-333333333333', 'editor', 'Content Editor', 'Manage content and master data')
ON CONFLICT (name) DO NOTHING;

-- 2. Seed Permissions
INSERT INTO permissions (name, resource, action, description) VALUES
('users:create', 'users', 'create', 'Create new users'),
('users:read', 'users', 'read', 'View user profiles'),
('users:update', 'users', 'update', 'Modify user details'),
('users:delete', 'users', 'delete', 'Remove users'),

('roles:read', 'roles', 'read', 'View roles and permissions'),
('roles:write', 'roles', 'write', 'Assign and update roles'),

('clients:create', 'clients', 'create', 'Add new clients'),
('clients:read', 'clients', 'read', 'View clients list and details'),
('clients:update', 'clients', 'update', 'Update client information'),
('clients:delete', 'clients', 'delete', 'Delete clients'),

('partners:create', 'partners', 'create', 'Add new partners'),
('partners:read', 'partners', 'read', 'View partners list and details'),
('partners:update', 'partners', 'update', 'Update partner information'),
('partners:delete', 'partners', 'delete', 'Delete partners'),

('solutions:create', 'solutions', 'create', 'Add solutions and features'),
('solutions:read', 'solutions', 'read', 'View solutions and features'),
('solutions:update', 'solutions', 'update', 'Update solutions'),
('solutions:delete', 'solutions', 'delete', 'Delete solutions'),

('team:create', 'team', 'create', 'Add team members'),
('team:read', 'team', 'read', 'View team members'),
('team:update', 'team', 'update', 'Update team members'),
('team:delete', 'team', 'delete', 'Delete team members'),

('contacts:read', 'contacts', 'read', 'View submitted contact messages'),
('contacts:update', 'contacts', 'update', 'Change contact message status'),
('contacts:delete', 'contacts', 'delete', 'Delete contact submissions'),

('company:update', 'company', 'update', 'Update company profile and offices')
ON CONFLICT (name) DO NOTHING;

-- 3. Assign All Permissions to super_admin
INSERT INTO role_permissions (role_id, permission_id)
SELECT '11111111-1111-1111-1111-111111111111', id FROM permissions
ON CONFLICT DO NOTHING;

-- 4. Seed Super Admin User (password: admin12345 using pgcrypto crypt or standard hash)
INSERT INTO users (id, email, username, password_hash, full_name, is_active)
VALUES (
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    'admin@company.com',
    'superadmin',
    crypt('admin12345', gen_salt('bf', 12)),
    'System Super Administrator',
    true
)
ON CONFLICT (email) DO NOTHING;

-- Assign super_admin role to the superadmin user
INSERT INTO user_roles (user_id, role_id)
VALUES (
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    '11111111-1111-1111-1111-111111111111'
)
ON CONFLICT DO NOTHING;

-- 5. Seed Company Profile
INSERT INTO company_profiles (id, company_name, tagline, about_text, vision, mission, founded_year, logo_url)
VALUES (
    'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
    'Unotek Digital Solusindo',
    'Empowering Enterprise Logistics & Supply Chain with Next-Gen Software',
    'Kami adalah software house spesialis enterprise solutions yang berfokus pada efisiensi supply chain, smart warehousing, dan automasi proses bisnis terintegrasi.',
    'Menjadi mitra transformasi digital nomor satu di Asia Tenggara untuk solusi industri & logistik terintegrasi.',
    'Mengembangkan produk perangkat lunak berkualitas tinggi, handal, dan berkinerja tinggi yang memberikan dampak bisnis nyata bagi setiap mitra dan klien.',
    2018,
    'https://assets.company.com/brand/unotek-logo.svg'
)
ON CONFLICT (id) DO NOTHING;

-- 6. Seed Office Locations
INSERT INTO office_locations (branch_name, branch_type, address, city, province, postal_code, country, phone, email, is_head_office, display_order)
VALUES
('Jakarta Head Office', 'head_office', 'Menara Standard Chartered Lt. 28, Jl. Prof. DR. Satrio No.164', 'Jakarta Selatan', 'DKI Jakarta', '12930', 'Indonesia', '+62 21 5550 1928', 'contact@company.com', true, 1),
('Bandung Tech Hub', 'development_center', 'Dago Digital Valley, Jl. Ir. H. Juanda No. 84', 'Bandung', 'Jawa Barat', '40132', 'Indonesia', '+62 22 7200 4811', 'dev@company.com', false, 2);

-- 7. Seed Master Clients
INSERT INTO clients (name, slug, industry, website_url, logo_url, description, is_featured, display_order)
VALUES
('PT Indomarco Prismatama (Indomaret)', 'indomaret', 'Retail & FMCG', 'https://indomaret.co.id', 'https://assets.company.com/clients/indomaret.svg', 'Implementasi sistem monitoring distribusi toko dan inventori regional.', true, 1),
('PT Sumber Alfaria Trijaya Tbk (Alfamart)', 'alfamart', 'Retail & Distribution', 'https://alfamart.co.id', 'https://assets.company.com/clients/alfamart.svg', 'Sistem integrasi warehouse batching and sorting.', true, 2),
('Pos Indonesia Logistics', 'pos-indonesia', 'Postal & Logistics', 'https://posindonesia.co.id', 'https://assets.company.com/clients/pos-indonesia.svg', 'Modernisasi track & trace dan warehouse hub nasional.', true, 3),
('DHL Supply Chain Indonesia', 'dhl-indonesia', 'Global Logistics', 'https://dhl.com', 'https://assets.company.com/clients/dhl.svg', 'Automasi manajemen cross-docking dan multi-client warehouse.', true, 4),
('Kalbe Farma Distribution', 'kalbe-farma', 'Pharmaceutical', 'https://kalbe.co.id', 'https://assets.company.com/clients/kalbe.svg', 'Cold chain warehouse compliance dan barcode validation system.', true, 5);

-- 8. Seed Master Product & Technology Partners
INSERT INTO product_partners (name, slug, partner_type, website_url, logo_url, description, display_order)
VALUES
('Zebra Technologies', 'zebra-technologies', 'hardware', 'https://zebra.com', 'https://assets.company.com/partners/zebra.svg', 'Penyedia barcode scanner, mobile computer enterprise, dan RFID hardware terkemuka.', 1),
('Honeywell Safety & Productivity', 'honeywell', 'hardware', 'https://honeywell.com', 'https://assets.company.com/partners/honeywell.svg', 'Partner resmi untuk supply chain mobility device dan industrial thermal printer.', 2),
('Amazon Web Services (AWS)', 'amazon-web-services', 'cloud', 'https://aws.amazon.com', 'https://assets.company.com/partners/aws.svg', 'Infrastruktur cloud enterprise dengan ketersediaan tinggi dan sertifikasi ISO 27001.', 3),
('SAP Integration Partner', 'sap', 'enterprise_software', 'https://sap.com', 'https://assets.company.com/partners/sap.svg', 'Konektivitas seamless dengan SAP S/4HANA & ECC untuk modul MM dan SD.', 4),
('Hikvision Industrial Automation', 'hikvision', 'iot_automation', 'https://hikvision.com', 'https://assets.company.com/partners/hikvision.svg', 'Partner solusi computer vision dan OCR gate system untuk pergudangan.', 5);

-- 9. Seed Solutions (Warehouse Management System & others)
INSERT INTO solutions (id, title, slug, tagline, summary, description, icon_url, hero_image_url, is_featured, display_order)
VALUES (
    'cccccccc-cccc-cccc-cccc-cccccccccccc',
    'Warehouse Management System (WMS)',
    'warehouse-management-system',
    'Smart, Scalable, and Real-Time Warehouse Control',
    'Sistem otomasi gudang modern yang mengoptimalkan seluruh pergerakan barang mulai dari receiving, putaway, picking, packing hingga dispatch secara real-time.',
    'WMS kami dirancang khusus untuk menangani kompleksitas pergudangan high-throughput dengan akurasi stok hingga 99.9%. Dilengkapi integrasi barcode, barcode scanner 2D, mobile RF, dan algoritma putaway/picking otomatis untuk memangkas waktu operasional hingga 45%.',
    'https://assets.company.com/icons/wms-box.svg',
    'https://assets.company.com/mockups/wms-hero-dashboard.png',
    true,
    1
),
(
    'dddddddd-dddd-dddd-dddd-dddddddddddd',
    'Transport Management System (TMS)',
    'transport-management-system',
    'End-to-End Fleet Optimization & Delivery Tracking',
    'Solusi manajemen armada pengiriman, optimalisasi rute cerdas (route planner), dan proof-of-delivery (e-POD) digital berbasis mobile app driver.',
    'Mengintegrasikan perencanaan armada dengan telematika GPS untuk meminimalkan biaya bahan bakar, mengontrol lead time pengiriman, dan memberikan tracking publik real-time kepada pelanggan.',
    'https://assets.company.com/icons/tms-truck.svg',
    'https://assets.company.com/mockups/tms-hero-map.png',
    true,
    2
)
ON CONFLICT (slug) DO NOTHING;

-- Seed Solution Features for WMS
INSERT INTO solution_features (solution_id, title, description, icon_url, display_order)
VALUES
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Real-Time Inventory Visibility', 'Pantau kuantiti dan lokasi stok di berbagai gudang hingga level bin/rak dengan pembaruan instan.', 'https://assets.company.com/icons/eye.svg', 1),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Intelligent Wave & Batch Picking', 'Algoritma rute picking terpendek untuk meningkatkan kecepatan pemenuhan order secara drastis.', 'https://assets.company.com/icons/route.svg', 2),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Multi-Warehouse & Multi-Client', 'Kelola banyak entitas gudang dan 3PL multi-tenancy dalam satu dashboard terpusat.', 'https://assets.company.com/icons/layers.svg', 3),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Mobile Handheld Scanner Ready', 'Aplikasi mobile Android responsif untuk scanning barcode 1D/2D dan QR Code secara offline/online.', 'https://assets.company.com/icons/scanner.svg', 4);

-- Seed Solution Modules for WMS
INSERT INTO solution_modules (solution_id, title, description, image_url, display_order)
VALUES
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Inbound & Putaway Automation', 'Penerimaan PO barang, validasi kuantiti, pencetakan barcode label otomatis, dan rekomendasi rak kosong.', 'https://assets.company.com/modules/inbound.png', 1),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Inventory & Cycle Count Control', 'Stock opname berkala tanpa perlu menghentikan operasional gudang, tracking lot number, dan tanggal kadaluarsa (FIFO/FEFO).', 'https://assets.company.com/modules/inventory.png', 2),
('cccccccc-cccc-cccc-cccc-cccccccccccc', 'Outbound & Dispatch Management', 'Verifikasi packing sebelum muat truk, integrasi kurir, dan cetak Surat Jalan otomatis.', 'https://assets.company.com/modules/outbound.png', 3);

-- 10. Seed Team Members
INSERT INTO team_members (full_name, position, department, bio, photo_url, linkedin_url, display_order, is_leadership)
VALUES
('Arya Wibowo', 'Chief Executive Officer', 'Executive', '15+ tahun pengalaman memimpin transformasi sistem digital supply chain di berbagai perusahaan multinasional.', 'https://assets.company.com/team/arya.jpg', 'https://linkedin.com/in/aryawibowo', 1, true),
('Siti Rahmawati', 'Chief Technology Officer', 'Technology', 'Mantan Principal Software Architect dengan spesialisasi distributed systems, high throughput event-driven architecture, dan cloud security.', 'https://assets.company.com/team/siti.jpg', 'https://linkedin.com/in/sitirahmawati', 2, true),
('Budi Santoso', 'Head of Product Development', 'Product', 'Berpengalaman mendesain produk enterprise B2B dan solusi IoT supply chain berorientasi efisiensi operasional.', 'https://assets.company.com/team/budi.jpg', 'https://linkedin.com/in/budisantoso', 3, true),
('Dimas Pratama', 'Lead Solutions Architect', 'Engineering', 'Spesialis integrasi ERP (SAP, Oracle, NetSuite) dan pengembangan core engine pergudangan modern.', 'https://assets.company.com/team/dimas.jpg', 'https://linkedin.com/in/dimaspratama', 4, false);
