CREATE TABLE locations
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    deleted_at DATETIME,
    name       TEXT     NOT NULL,
    address    TEXT,
    lat        REAL,
    lon        REAL
);

CREATE INDEX idx_locations_deleted_at ON locations (deleted_at);

CREATE TABLE projects
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at  DATETIME NOT NULL DEFAULT current_timestamp,
    deleted_at  DATETIME,
    slug        TEXT     NOT NULL,
    name        TEXT     NOT NULL,
    location_id INTEGER
        CONSTRAINT fk_projects_location REFERENCES locations (id)
);

CREATE INDEX idx_projects_deleted_at ON projects (deleted_at);

CREATE TABLE versions
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at  DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at  DATETIME NOT NULL DEFAULT current_timestamp,
    deleted_at  DATETIME,
    name        TEXT     NOT NULL,
    description TEXT,
    project_id  INTEGER
        CONSTRAINT fk_version_project REFERENCES projects (id)
);

CREATE INDEX idx_versions_deleted_at ON versions (deleted_at);

CREATE TABLE files
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    deleted_at DATETIME,
    name       TEXT     NOT NULL,
    size       INTEGER  NOT NULL,
    path       TEXT     NOT NULL
);

CREATE INDEX idx_files_deleted_at ON files (deleted_at);

CREATE TABLE versions_files
(
    version_id INTEGER NOT NULL,
    file_id    INTEGER NOT NULL,
    PRIMARY KEY (version_id, file_id),
    CONSTRAINT fk_versions_files_version FOREIGN KEY (version_id) REFERENCES versions (id),
    CONSTRAINT fk_versions_files_file FOREIGN KEY (file_id) REFERENCES files (id)
);
