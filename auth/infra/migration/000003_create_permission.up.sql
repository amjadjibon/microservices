-- Create the UserPermission table to manage user permissions
CREATE TABLE IF NOT EXISTS auth_permission
(
    id SERIAL,
    user_id INT NOT NULL,
    permission_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (user_id) REFERENCES auth_user (id),
    PRIMARY KEY (id)
);
