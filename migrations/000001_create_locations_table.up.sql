CREATE TABLE locations
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT      NOT NULL,
    address    TEXT,
    lat        REAL,
    lon        REAL
);
