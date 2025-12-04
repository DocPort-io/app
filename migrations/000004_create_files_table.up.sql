CREATE TABLE files
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    name       TEXT     NOT NULL,
    size       INTEGER  NOT NULL,
    path       TEXT     NOT NULL
);
