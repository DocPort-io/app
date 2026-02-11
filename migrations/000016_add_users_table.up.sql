CREATE TABLE users
(
    id             BIGSERIAL PRIMARY KEY,
    created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name           TEXT      NOT NULL,
    email          TEXT      NOT NULL UNIQUE,
    email_verified BOOLEAN   NOT NULL DEFAULT FALSE
);
