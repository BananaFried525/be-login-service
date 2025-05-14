CREATE TYPE user_status AS ENUM (
    'ACTIVE',
    'INACTIVE',
    'DELETED',
    'SUSPENDED'
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    referral_code VARCHAR(12),
    user_name VARCHAR(50),
    email TEXT NOT NULL,
    status user_status NOT NULL DEFAULT 'ACTIVE',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE(email);

ALTER TABLE users ADD CONSTRAINT unique_referral_code UNIQUE(referral_code);