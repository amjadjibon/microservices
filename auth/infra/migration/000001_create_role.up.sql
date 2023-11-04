-- Create the Role table to manage user roles and permissions
CREATE TABLE IF NOT EXISTS auth_role
(
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id)
);
