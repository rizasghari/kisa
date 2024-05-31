-- Create Users Table
CREATE TABLE users
(
    id            UUID PRIMARY KEY             DEFAULT gen_random_uuid(),
    username      VARCHAR(50) UNIQUE  NOT NULL,
    email         VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255)        NOT NULL,
    created_at    TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create URLs Table
CREATE TABLE urls
(
    id           UUID PRIMARY KEY            DEFAULT gen_random_uuid(),
    user_id      UUID REFERENCES users (id),
    original_url TEXT               NOT NULL,
    short_url    VARCHAR(10) UNIQUE NOT NULL,
    created_at   TIMESTAMP          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    access_count INTEGER            NOT NULL DEFAULT 0
);

-- Create URL Access Logs Table
CREATE TABLE url_access_logs
(
    id          UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    url_id      UUID REFERENCES urls (id),
    accessed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    referrer    TEXT,
    user_agent  TEXT
);
