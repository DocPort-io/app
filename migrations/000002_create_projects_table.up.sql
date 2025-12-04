CREATE TABLE projects
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at  DATETIME NOT NULL DEFAULT current_timestamp,
    slug        TEXT     NOT NULL,
    name        TEXT     NOT NULL,
    location_id INTEGER
        CONSTRAINT fk_projects_location REFERENCES locations (id)
);
