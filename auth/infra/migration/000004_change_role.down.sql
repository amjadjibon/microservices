-- Drop auth role table name column unique
ALTER TABLE auth_role
    DROP CONSTRAINT auth_role_name_unique;
