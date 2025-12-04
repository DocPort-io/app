-- Projects

-- name: GetProject :one
SELECT *
FROM projects
WHERE id = ?
LIMIT 1;

-- name: ListProjectsWithLocations :many
SELECT projects.*,
       locations.name AS location_name,
       locations.lat  AS location_lat,
       locations.lon  AS location_lon
FROM projects
         LEFT JOIN locations ON locations.id = projects.location_id
ORDER BY projects.created_at DESC;

-- name: CreateProject :one
INSERT INTO projects (slug, name, location_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET slug        = ?,
    name        = ?,
    location_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = ?;

-- Versions

-- name: ListVersions :many
SELECT *
FROM versions
ORDER BY created_at DESC;

-- name: ListVersionsByProjectId :many
SELECT *
FROM versions
WHERE project_id = ?
ORDER BY created_at DESC;

-- name: GetVersion :one
SELECT *
FROM versions
WHERE id = ?
LIMIT 1;

-- name: CreateVersion :one
INSERT INTO versions (name, description, project_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateVersion :one
UPDATE versions
SET name        = ?,
    description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteVersion :exec
DELETE
FROM versions
WHERE id = ?;

-- Locations

-- name: ListLocations :many
SELECT *
FROM locations
ORDER BY created_at DESC;

-- Files

-- name: ListFiles :many
SELECT *
FROM files
ORDER BY created_at DESC;

-- name: ListFilesByVersionId :many
SELECT files.*
FROM files
         INNER JOIN versions_files ON files.id = versions_files.file_id
WHERE versions_files.version_id = ?
ORDER BY files.created_at DESC;

-- name: GetFile :one
SELECT *
FROM files
WHERE id = ?
LIMIT 1;

-- name: CreateFile :one
INSERT INTO files (name, size, path)
VALUES (?, ?, ?)
RETURNING *;

-- name: DeleteFile :exec
DELETE
FROM files
WHERE id = ?;

-- name: AttachFileToVersion :exec
INSERT INTO versions_files (version_id, file_id)
VALUES (?, ?);