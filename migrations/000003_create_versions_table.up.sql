CREATE TABLE versions
(
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name        TEXT      NOT NULL,
    description TEXT,
    project_id  INTEGER   NOT NULL
        CONSTRAINT fk_version_project REFERENCES projects (id) ON DELETE CASCADE
);
