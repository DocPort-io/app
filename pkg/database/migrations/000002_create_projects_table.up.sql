CREATE TABLE projects
(
    id          BIGSERIAL PRIMARY KEY,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    slug        TEXT      NOT NULL,
    name        TEXT      NOT NULL,
    location_id INTEGER
        CONSTRAINT fk_projects_location REFERENCES locations (id)
);
