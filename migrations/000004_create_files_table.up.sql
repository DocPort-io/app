CREATE TABLE files
(
    id         BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT      NOT NULL,
    size       INTEGER   NOT NULL,
    path       TEXT      NOT NULL
);
