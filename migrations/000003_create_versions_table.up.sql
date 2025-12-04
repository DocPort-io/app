CREATE TABLE versions
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at  DATETIME NOT NULL DEFAULT current_timestamp,
    name        TEXT     NOT NULL,
    description TEXT,
    project_id  INTEGER  NOT NULL
        CONSTRAINT fk_version_project REFERENCES projects (id) ON DELETE CASCADE
);
