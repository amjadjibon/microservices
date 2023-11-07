-- Change auth role table name column unique
ALTER TABLE auth_role
    ADD CONSTRAINT auth_role_name_unique UNIQUE (name);
