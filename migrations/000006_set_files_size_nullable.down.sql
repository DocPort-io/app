CREATE TABLE files_dg_tmp
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    name       TEXT     NOT NULL,
    size       INTEGER  NOT NULL,
    path       TEXT     NOT NULL
);

INSERT INTO files_dg_tmp(id, created_at, updated_at, name, size, path)
SELECT id, created_at, updated_at, name, size, path
FROM files;

DROP TABLE files;

ALTER TABLE files_dg_tmp
    RENAME TO files;
