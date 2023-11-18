CREATE TABLE title_title (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    year INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE title_language (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE title_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE title_content (
    id SERIAL PRIMARY KEY,
    title_id INTEGER REFERENCES title_title(id),
    name VARCHAR(255) NOT NULL,
    type_id INTEGER REFERENCES title_type(id),
    year INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    
    CONSTRAINT fk_content_title FOREIGN KEY (title_id) REFERENCES title_title(id),
    CONSTRAINT fk_content_type FOREIGN KEY (type_id) REFERENCES title_type(id)
);

CREATE TABLE title_content_language (
    content_id INTEGER REFERENCES title_content(id),
    language_id INTEGER REFERENCES title_language(id),
    PRIMARY KEY (content_id, language_id),
    CONSTRAINT fk_content_language_content FOREIGN KEY (content_id) REFERENCES title_content(id),
    CONSTRAINT fk_content_language_language FOREIGN KEY (language_id) REFERENCES title_language(id)
);