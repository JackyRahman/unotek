# Unotek Enterprise Backend

Backend API untuk website profil perusahaan software house (**Unotek**) yang dibangun menggunakan **Go (Golang)** dengan **Clean Architecture**, **PostgreSQL**, **Role-Based Access Control (RBAC)**, **Swagger OpenAPI**, dan standar keamanan enterprise (*zero known vulnerabilities*).

---

## 🚀 Fitur Utama

- **Go Standard Clean Architecture**: Pemisahan lapisan yang jelas antara *Domain Entities*, *Use Cases / Business Logic*, *Repositories*, dan *HTTP Delivery*.
- **PostgreSQL Database Design**:
  - Standar penamaan tabel profesional (*plural, snake_case*).
  - Primary Key UUID v4 (`gen_random_uuid()`).
  - Audit timestamps (`created_at`, `updated_at`, `deleted_at` untuk soft delete).
- **Role-Based Access Control (RBAC)**:
  - Role bertingkat: `super_admin`, `admin`, `editor`.
  - Granular permissions per resource (`clients:read`, `clients:create`, `contacts:update`, dll).
- **Security & Best Practices**:
  - Password hashing menggunakan **Bcrypt** (cost 12).
  - Stateless **JWT Authentication** (JSON Web Token HS256).
  - Perlindungan SQL Injection (GORM parameterized query).
  - Middleware: CORS, In-Memory Rate Limiting, Request ID tracing, Panic Recovery, dan Validasi Input Request DTO.
- **Master Data Lengkap**:
  - 🏢 **Company Profile & Offices**: Data profil perusahaan dan multi-lokasi kantor (Jakarta HQ, Bandung Tech Hub, dll).
  - 🤝 **Clients**: Master data klien & portofolio enterprise.
  - 🔌 **Product & Technology Partners**: Master mitra teknologi, hardware (Zebra, Honeywell), cloud (AWS), dan ERP (SAP).
  - 📦 **Solutions**: Master katalog solusi (WMS - *Warehouse Management System*, TMS, dll) lengkap dengan fitur modular dan sub-modul.
  - 👥 **Team Members**: Profil tim manajemen dan direksi.
  - 📩 **Contact Submissions (Leads)**: Formulir konsultasi & prospek dari landing page publik.
- **Interactive Swagger Documentation**: Akses UI dokumentasi interaktif lengkap di `/swagger/index.html`.
- **Automated Initial Seeder**: Menyiapkan akun Super Admin awal dan data master saat pertama kali dijalankan.

---

## 📋 Prasyarat Sistem (Prerequisites)

- **Go**: Versi **1.22** atau lebih baru (`go version`)
- **PostgreSQL**: Versi **14** / **15** / **16**
- **Git**: Versi 2.x
- **Docker & Docker Compose** *(Opsional, jika ingin menjalankan via kontainer)*

---

## 🛠️ Panduan Menyiapkan Git Repository

Jalankan perintah berikut di direktori root project untuk menginisialisasi repo dan push ke GitHub:

```bash
# Inisialisasi git repository
git init

# Tambahkan remote repository
git remote add origin https://github.com/JackyRahman/unotek.git

# Buat branch utama 'main'
git branch -M main

# Stage semua file dan lakukan first commit
git add .
git commit -m "feat: initial commit enterprise clean architecture backend with rbac, migrations, seeders, and swagger"

# Push ke repository remote
git push -u origin main
```

---

## ⚙️ Konfigurasi Environment (`.env`)

Salin file `.env.example` menjadi `.env` lalu sesuaikan konfigurasi database Anda:

```bash
cp .env.example .env
```

Isi konfigurasi default:
```env
APP_NAME=unotek-backend
APP_ENV=development
APP_PORT=8080
APP_DEBUG=true
APP_URL=http://localhost:8080

# Database Configuration (PostgreSQL)
# Format 1 Baris Connection URL (Recommended):
DATABASE_URL=postgres://postgres:admin12345@100.xxx.247.22:5432/itsm_dev?sslmode=disable

# Atau Format Parameter Terpisah (Fallback):
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=postgres
# DB_NAME=unotek_company_db
# DB_SSLMODE=disable
# DB_TIMEZONE=Asia/Jakarta

# JWT Security
JWT_SECRET=super-secret-key-change-this-in-production-min-32-chars-long!
JWT_EXPIRATION_HOURS=24
JWT_REFRESH_EXPIRATION_DAYS=7

# Rate Limiter
RATE_LIMIT_RPS=20
RATE_LIMIT_BURST=40

```

---

## 💾 Perintah Database Migration & Seeder

Migration dan seeder dipisahkan dari server runtime agar server dapat langsung menyala secara instan tanpa query overhead:

```bash
# 1. Jalankan migrasi tabel saja
go run ./cmd/migrate/main.go up

# 2. Jalankan seeder master data & Super Admin
go run ./cmd/migrate/main.go seed

# 3. Jalankan migrasi sekaligus seeder
go run ./cmd/migrate/main.go all

# 4. Rollback / drop semua tabel (reset)
go run ./cmd/migrate/main.go drop
```

---

## 🏃 Cara Menjalankan Server API

Server API sekarang murni menjalankan HTTP engine tanpa overhead migration:

```bash
# Jalankan server API (instan dalam milidetik)
go run ./cmd/api/main.go
```

---

### Opsi 2: Menjalankan via Docker Compose (Paling Praktis)

Cukup jalankan satu perintah:
```bash
docker compose up -d
```
Docker Compose akan secara otomatis menjalankan PostgreSQL dan aplikasi backend sekaligus.

---

## 🔑 Kredensial Super Admin Awal (Seeder)

Setelah aplikasi dijalankan pertama kali, akun Super Admin otomatis tersedia untuk pengujian:

| Field | Nilai |
|---|---|
| **Email** | `admin@company.com` |
| **Username** | `superadmin` |
| **Password** | `admin12345` |
| **Role** | `super_admin` |

---

## 📖 Dokumentasi API & Swagger UI

Buka browser dan akses alamat berikut saat server berjalan:

👉 **[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)**

Di Swagger UI, Anda dapat langsung mencoba:
1. Klik endpoint `/api/v1/auth/login` $\rightarrow$ klik **Try it out** $\rightarrow$ masukkan kredensial `admin@company.com` / `admin12345`.
2. Salin nilai `access_token` dari respons JSON.
3. Klik tombol **Authorize** di pojok kanan atas Swagger, ketik `Bearer <token_anda>`, lalu klik **Authorize**.
4. Semua endpoint Admin (`/api/v1/admin/*`) kini siap diuji langsung dari browser!

---

## 📡 Daftar Endpoint Master Data

### 1. Autentikasi
| Method | Endpoint | Keterangan | Akses |
|---|---|---|---|
| `POST` | `/api/v1/auth/login` | Login user & generate JWT Token | Publik |
| `GET` | `/api/v1/auth/profile` | Ambil data profil & permission user yang login | Bearer Token |

### 2. Publik (Landing Page / Frontend)
| Method | Endpoint | Keterangan | Akses |
|---|---|---|---|
| `GET` | `/health` | Healthcheck server status | Publik |
| `GET` | `/api/v1/company/profile` | Profil perusahaan, visi, misi, tagline | Publik |
| `GET` | `/api/v1/company/offices` | Daftar kantor cabang & alamat | Publik |
| `GET` | `/api/v1/clients` | List klien terdaftar & logo portofolio | Publik |
| `GET` | `/api/v1/clients/:slug` | Detail klien berdasarkan slug | Publik |
| `GET` | `/api/v1/partners` | List partner teknologi & hardware | Publik |
| `GET` | `/api/v1/partners/:slug` | Detail partner berdasarkan slug | Publik |
| `GET` | `/api/v1/solutions` | List katalog produk/solusi (WMS, TMS) | Publik |
| `GET` | `/api/v1/solutions/:slug` | Detail solusi lengkap dengan fitur & modul | Publik |
| `GET` | `/api/v1/team` | List tim manajemen & direksi | Publik |
| `POST` | `/api/v1/contact` | Kirim pesan / lead / form kontak dari web | Publik |

### 3. CMS & Admin Panel (Guarded by JWT & Granular RBAC)
| Method | Endpoint | Permission Diperlukan |
|---|---|---|
| `GET` | `/api/v1/admin/clients` | `clients:read` |
| `POST` | `/api/v1/admin/clients` | `clients:create` |
| `GET` | `/api/v1/admin/clients/:id` | `clients:read` |
| `PUT` | `/api/v1/admin/clients/:id` | `clients:update` |
| `DELETE`| `/api/v1/admin/clients/:id` | `clients:delete` |
| `GET` | `/api/v1/admin/partners` | `partners:read` |
| `POST` | `/api/v1/admin/partners` | `partners:create` |
| `PUT` | `/api/v1/admin/partners/:id` | `partners:update` |
| `DELETE`| `/api/v1/admin/partners/:id` | `partners:delete` |
| `GET` | `/api/v1/admin/solutions` | `solutions:read` |
| `POST` | `/api/v1/admin/solutions` | `solutions:create` |
| `PUT` | `/api/v1/admin/solutions/:id` | `solutions:update` |
| `DELETE`| `/api/v1/admin/solutions/:id` | `solutions:delete` |
| `POST` | `/api/v1/admin/solutions/:id/features` | `solutions:update` |
| `DELETE`| `/api/v1/admin/solutions/:id/features/:f_id`| `solutions:update` |
| `POST` | `/api/v1/admin/solutions/:id/modules` | `solutions:update` |
| `DELETE`| `/api/v1/admin/solutions/:id/modules/:m_id` | `solutions:update` |
| `GET` | `/api/v1/admin/team` | `team:read` |
| `POST` | `/api/v1/admin/team` | `team:create` |
| `PUT` | `/api/v1/admin/team/:id` | `team:update` |
| `DELETE`| `/api/v1/admin/team/:id` | `team:delete` |
| `GET` | `/api/v1/admin/contacts` | `contacts:read` |
| `PATCH` | `/api/v1/admin/contacts/:id/status` | `contacts:update` |
| `DELETE`| `/api/v1/admin/contacts/:id` | `contacts:delete` |
| `PUT` | `/api/v1/admin/company/profile` | `company:update` |
| `POST` | `/api/v1/admin/company/offices` | `company:update` |
| `GET` | `/api/v1/admin/rbac/roles` | `super_admin` role |
| `GET` | `/api/v1/admin/rbac/permissions` | `super_admin` role |

---

## 🧪 Contoh Pengujian via cURL

### 1. Login Admin
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "admin@company.com",
    "password": "admin12345"
  }'
```

### 2. Mengambil Profil User yang Sedang Login
```bash
curl -X GET http://localhost:8080/api/v1/auth/profile \
  -H "Authorization: Bearer <TOKEN_HASIL_LOGIN>"
```

### 3. Mengambil Daftar Klien (Publik)
```bash
curl -X GET "http://localhost:8080/api/v1/clients?page=1&per_page=10"
```

### 4. Mengirim Pesan Kontak (Publik)
```bash
curl -X POST http://localhost:8080/api/v1/contact \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Rian Kusuma",
    "email": "rian@perusahaan.com",
    "phone": "+6281298765432",
    "company_name": "PT Sumber Distribusi",
    "service_interest": "Warehouse Management System (WMS)",
    "subject": "Konsultasi Implementasi Gudang Dingin",
    "message": "Halo, kami ingin mengetahui estimasi waktu implementasi WMS untuk 2 gudang kami."
  }'
```

---

## 🏛️ Struktur Folder (Clean Architecture)

```
unotek/
├── cmd/
│   └── api/
│       └── main.go                  # Main application entry point & graceful shutdown
├── docs/                            # Swagger 2.0 & OpenAPI specifications
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── config/                      # Environment variables & application configs
│   │   └── config.go
│   ├── delivery/
│   │   └── http/
│   │       ├── dto/                 # Request & Response Data Transfer Objects
│   │       ├── handler/             # HTTP Controllers / Request Handlers
│   │       ├── middleware/          # JWT, RBAC Guard, Rate Limiter, CORS, Request ID
│   │       └── router/              # Route mapping and engine setup
│   ├── domain/                      # Enterprise entities, interfaces, & contracts
│   │   ├── audit.go
│   │   ├── auth.go
│   │   ├── client.go
│   │   ├── company.go
│   │   ├── contact.go
│   │   ├── partner.go
│   │   ├── rbac.go
│   │   ├── solution.go
│   │   ├── team.go
│   │   └── user.go
│   ├── repository/
│   │   └── postgres/                # PostgreSQL persistence layer & GORM queries
│   │       ├── client_repository.go
│   │       ├── company_repository.go
│   │       ├── contact_repository.go
│   │       ├── partner_repository.go
│   │       ├── rbac_repository.go
│   │       ├── seeder.go            # Comprehensive master data seeder
│   │       ├── solution_repository.go
│   │       ├── team_repository.go
│   │       └── user_repository.go
│   └── usecase/                     # Pure business logic orchestrators
│       ├── auth_usecase.go
│       ├── client_usecase.go
│       ├── company_usecase.go
│       ├── contact_usecase.go
│       ├── partner_usecase.go
│       ├── rbac_usecase.go
│       ├── solution_usecase.go
│       └── team_usecase.go
├── migrations/                      # PostgreSQL DDL migrations (.up.sql & .down.sql)
│   ├── 000001_create_rbac_and_users.up.sql
│   ├── 000001_create_rbac_and_users.down.sql
│   ├── 000002_create_masters_and_content.up.sql
│   ├── 000002_create_masters_and_content.down.sql
│   └── 000003_seed_initial_data.up.sql
├── pkg/                             # Reusable packages & utilities
│   ├── database/                    # GORM PostgreSQL pool setup
│   ├── errors/                      # Standard application error definitions
│   ├── pagination/                  # Universal pagination helpers
│   ├── response/                    # Standard JSON response envelope
│   └── security/                    # Bcrypt hash & JWT signing/validation
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 📄 Lisensi

Proyek ini berada di bawah lisensi [MIT](LICENSE).
