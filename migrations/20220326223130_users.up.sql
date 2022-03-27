CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    uuid UUID DEFAULT uuid_generate_v4(),
    login VARCHAR(15) NOT NULL,
    email VARCHAR(50) NOT NULL,
    hash VARCHAR(255) NOT NULL,
    role INT DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    UNIQUE (login, email),
    UNIQUE (uuid)
);