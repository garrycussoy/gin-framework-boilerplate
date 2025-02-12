CREATE TABLE IF NOT EXISTS "user" (
    id uuid PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    branchId VARCHAR(4),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255),
    updated_at TIMESTAMP,
    updated_by uuid,
    custom_1 JSONB,
    custom_2 JSONB,
    custom_3 JSONB
);
