CREATE TABLE tracker.categories (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL CHECK (char_length(name) BETWEEN 3 AND 100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT categories_user_id_name_key UNIQUE (user_id, name)
);

