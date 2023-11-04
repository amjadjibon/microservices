-- Create the User table to store user information
CREATE TABLE IF NOT EXISTS auth_user
(
    id SERIAL,
    username VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    is_verified BOOLEAN NOT NULL,
    gender VARCHAR(10),
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (role_id) REFERENCES auth_role (id),
    PRIMARY KEY (id)
);
