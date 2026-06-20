CREATE TABLE tracker.expenses (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES tracker.users(id) ON DELETE CASCADE,
    category_id uuid NOT NULL REFERENCES tracker.categories(id) ON DELETE RESTRICT,
    amount NUMERIC(12,2) NOT NULL CHECK (amount >= 0),
    description TEXT,
    expense_date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    version INTEGER DEFAULT 1
);
