-- Table: company_profiles
CREATE TABLE IF NOT EXISTS company_profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    company_name VARCHAR(150) NOT NULL,
    tagline VARCHAR(255),
    about_text TEXT,
    vision TEXT,
    mission TEXT,
    founded_year INT,
    logo_url TEXT,
    favicon_url TEXT,
    meta_title VARCHAR(255),
    meta_description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: office_locations
CREATE TABLE IF NOT EXISTS office_locations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    branch_name VARCHAR(100) NOT NULL,
    branch_type VARCHAR(50) NOT NULL DEFAULT 'head_office',
    address TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    province VARCHAR(100),
    postal_code VARCHAR(20),
    country VARCHAR(100) NOT NULL DEFAULT 'Indonesia',
    phone VARCHAR(50),
    email VARCHAR(150),
    working_hours VARCHAR(100),
    google_maps_url TEXT,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    is_head_office BOOLEAN NOT NULL DEFAULT false,
    display_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_offices_active ON office_locations(is_active);
CREATE INDEX IF NOT EXISTS idx_offices_order ON office_locations(display_order);

-- Table: clients
CREATE TABLE IF NOT EXISTS clients (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    slug VARCHAR(180) NOT NULL UNIQUE,
    industry VARCHAR(100),
    website_url TEXT,
    logo_url TEXT NOT NULL,
    description TEXT,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    display_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_clients_slug ON clients(slug);
CREATE INDEX IF NOT EXISTS idx_clients_featured ON clients(is_featured);
CREATE INDEX IF NOT EXISTS idx_clients_order ON clients(display_order);
CREATE INDEX IF NOT EXISTS idx_clients_deleted_at ON clients(deleted_at);

-- Table: product_partners
CREATE TABLE IF NOT EXISTS product_partners (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) NOT NULL,
    slug VARCHAR(180) NOT NULL UNIQUE,
    partner_type VARCHAR(50) NOT NULL DEFAULT 'technology',
    website_url TEXT,
    logo_url TEXT NOT NULL,
    description TEXT,
    display_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_partners_slug ON product_partners(slug);
CREATE INDEX IF NOT EXISTS idx_partners_type ON product_partners(partner_type);
CREATE INDEX IF NOT EXISTS idx_partners_order ON product_partners(display_order);
CREATE INDEX IF NOT EXISTS idx_partners_deleted_at ON product_partners(deleted_at);

-- Table: solutions
CREATE TABLE IF NOT EXISTS solutions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(150) NOT NULL,
    slug VARCHAR(180) NOT NULL UNIQUE,
    tagline VARCHAR(255),
    summary TEXT,
    description TEXT,
    icon_url TEXT,
    hero_image_url TEXT,
    brochure_url TEXT,
    display_order INT NOT NULL DEFAULT 0,
    is_featured BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_solutions_slug ON solutions(slug);
CREATE INDEX IF NOT EXISTS idx_solutions_featured ON solutions(is_featured);
CREATE INDEX IF NOT EXISTS idx_solutions_order ON solutions(display_order);
CREATE INDEX IF NOT EXISTS idx_solutions_deleted_at ON solutions(deleted_at);

-- Table: solution_features
CREATE TABLE IF NOT EXISTS solution_features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    solution_id UUID NOT NULL REFERENCES solutions(id) ON DELETE CASCADE,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    icon_url TEXT,
    display_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_features_solution_id ON solution_features(solution_id);
CREATE INDEX IF NOT EXISTS idx_features_order ON solution_features(display_order);

-- Table: solution_modules
CREATE TABLE IF NOT EXISTS solution_modules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    solution_id UUID NOT NULL REFERENCES solutions(id) ON DELETE CASCADE,
    title VARCHAR(150) NOT NULL,
    description TEXT,
    image_url TEXT,
    display_order INT NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_modules_solution_id ON solution_modules(solution_id);
CREATE INDEX IF NOT EXISTS idx_modules_order ON solution_modules(display_order);

-- Table: team_members
CREATE TABLE IF NOT EXISTS team_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(150) NOT NULL,
    position VARCHAR(100) NOT NULL,
    department VARCHAR(100),
    bio TEXT,
    photo_url TEXT,
    linkedin_url TEXT,
    email VARCHAR(150),
    display_order INT NOT NULL DEFAULT 0,
    is_leadership BOOLEAN NOT NULL DEFAULT false,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_team_order ON team_members(display_order);
CREATE INDEX IF NOT EXISTS idx_team_leadership ON team_members(is_leadership);
CREATE INDEX IF NOT EXISTS idx_team_deleted_at ON team_members(deleted_at);

-- Table: contact_submissions (Leads / Inquiries)
CREATE TABLE IF NOT EXISTS contact_submissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR(150) NOT NULL,
    email VARCHAR(150) NOT NULL,
    phone VARCHAR(50),
    company_name VARCHAR(150),
    service_interest VARCHAR(100),
    subject VARCHAR(200) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'new',
    notes TEXT,
    ip_address VARCHAR(45),
    user_agent TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_contacts_status ON contact_submissions(status);
CREATE INDEX IF NOT EXISTS idx_contacts_email ON contact_submissions(email);
CREATE INDEX IF NOT EXISTS idx_contacts_created_at ON contact_submissions(created_at);
CREATE INDEX IF NOT EXISTS idx_contacts_deleted_at ON contact_submissions(deleted_at);
