CREATE TABLE external_auth
(
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id     BIGINT    NOT NULL
        CONSTRAINT fk_external_auth_user REFERENCES users (id) ON DELETE CASCADE,
    provider    TEXT      NOT NULL,
    provider_id TEXT      NOT NULL,
    UNIQUE (user_id, provider)
);

CREATE INDEX idx_external_auth_user_id ON external_auth (user_id);
