-- Projects

-- name: GetProject :one
SELECT id,
       created_at,
       updated_at,
       slug,
       name,
       location_id
FROM projects
WHERE id = ?
LIMIT 1;

-- name: CountProjects :one
SELECT count(projects.id)
FROM projects;

-- name: ListProjectsWithLocations :many
SELECT projects.id,
       projects.created_at,
       projects.updated_at,
       projects.slug,
       projects.name,
       projects.location_id,
       locations.name AS location_name,
       locations.lat  AS location_lat,
       locations.lon  AS location_lon
FROM projects
         LEFT JOIN locations ON locations.id = projects.location_id
ORDER BY projects.created_at DESC
LIMIT ? OFFSET ?;

-- name: CreateProject :one
INSERT INTO projects (slug, name, location_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET updated_at  = current_timestamp,
    slug        = ?,
    name        = ?,
    location_id = ?
WHERE id = ?
RETURNING *;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = ?;

-- Versions

-- name: CountVersionsByProjectId :one
SELECT count(versions.id)
FROM versions
WHERE project_id = ?;

-- name: ListVersionsByProjectId :many
SELECT *
FROM versions
WHERE project_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetVersion :one
SELECT id,
       created_at,
       updated_at,
       name,
       description,
       project_id
FROM versions
WHERE id = ?
LIMIT 1;

-- name: CreateVersion :one
INSERT INTO versions (name, description, project_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: UpdateVersion :one
UPDATE versions
SET updated_at  = current_timestamp,
    name        = ?,
    description = ?
WHERE id = ?
RETURNING *;

-- name: DeleteVersion :exec
DELETE
FROM versions
WHERE id = ?;

-- Locations

-- name: ListLocations :many
SELECT id,
       created_at,
       updated_at,
       name,
       address,
       lat,
       lon
FROM locations
ORDER BY created_at DESC;

-- Files

-- name: CountFilesByVersionId :one
SELECT count(files.id)
FROM files
         INNER JOIN versions_files ON files.id = versions_files.file_id
WHERE versions_files.version_id = ?;

-- name: ListFilesByVersionId :many
SELECT files.id,
       files.created_at,
       files.updated_at,
       files.name,
       files.size,
       files.path,
       files.mime_type,
       files.is_complete
FROM files
         INNER JOIN versions_files ON files.id = versions_files.file_id
WHERE versions_files.version_id = ?
ORDER BY files.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetFile :one
SELECT id,
       created_at,
       updated_at,
       name,
       size,
       path,
       mime_type,
       is_complete
FROM files
WHERE id = ?
LIMIT 1;

-- name: CreateFile :one
INSERT INTO files (name)
VALUES (?)
RETURNING *;

-- name: UpdateFileWithUploadedFile :one
UPDATE files
SET updated_at  = current_timestamp,
    size        = ?,
    path        = ?,
    mime_type   = ?,
    is_complete = TRUE
WHERE id = ?
  AND is_complete = FALSE
RETURNING *;

-- name: DeleteFile :exec
DELETE
FROM files
WHERE id = ?;

-- name: AttachFileToVersion :exec
INSERT INTO versions_files (version_id, file_id)
VALUES (?, ?);

-- name: DetachFileFromVersion :exec
DELETE
FROM versions_files
WHERE version_id = ?
  AND file_id = ?;
