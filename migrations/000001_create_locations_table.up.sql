CREATE TABLE locations
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at DATETIME NOT NULL DEFAULT current_timestamp,
    updated_at DATETIME NOT NULL DEFAULT current_timestamp,
    name       TEXT     NOT NULL,
    address    TEXT,
    lat        REAL,
    lon        REAL
);
