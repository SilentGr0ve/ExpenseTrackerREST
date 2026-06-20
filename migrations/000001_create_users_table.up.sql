CREATE SCHEMA tracker;

CREATE TABLE tracker.users (
    id uuid PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    version INTEGER DEFAULT 1
);




