-- Projects

-- name: GetProject :one
SELECT id,
       created_at,
       updated_at,
       slug,
       name,
       location_id
FROM projects
WHERE id = $1
LIMIT 1;

-- name: CountProjects :one
SELECT count(projects.id)
FROM projects;

-- name: ListProjects :many
SELECT id,
       created_at,
       updated_at,
       slug,
       name,
       location_id
FROM projects
ORDER BY created_at DESC
LIMIT sqlc.arg('limit')::BIGINT OFFSET sqlc.arg('offset')::BIGINT;

-- name: CreateProject :one
INSERT INTO projects (slug, name, location_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET updated_at  = CURRENT_TIMESTAMP,
    slug        = $2,
    name        = $3,
    location_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE
FROM projects
WHERE id = $1;

-- Versions

-- name: CountVersionsByProjectId :one
SELECT count(versions.id)
FROM versions
WHERE project_id = $1;

-- name: ListVersions :many
SELECT *
FROM versions
WHERE (sqlc.narg('projectId')::BIGINT IS NULL OR project_id = sqlc.narg('projectId'))
ORDER BY created_at DESC
LIMIT sqlc.arg('limit')::BIGINT OFFSET sqlc.arg('offset')::BIGINT;

-- name: GetVersion :one
SELECT id,
       created_at,
       updated_at,
       name,
       description,
       project_id
FROM versions
WHERE id = $1
LIMIT 1;

-- name: CreateVersion :one
INSERT INTO versions (name, description, project_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateVersion :one
UPDATE versions
SET updated_at  = CURRENT_TIMESTAMP,
    name        = $2,
    description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteVersion :exec
DELETE
FROM versions
WHERE id = $1;

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
WHERE versions_files.version_id = $1;

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
WHERE versions_files.version_id = $1
ORDER BY files.created_at DESC
LIMIT sqlc.arg('limit')::BIGINT OFFSET sqlc.arg('offset')::BIGINT;

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
WHERE id = $1
LIMIT 1;

-- name: CreateFile :one
INSERT INTO files (name)
VALUES ($1)
RETURNING *;

-- name: UpdateFileWithUploadedFile :one
UPDATE files
SET updated_at  = current_timestamp,
    size        = $2,
    path        = $3,
    mime_type   = $4,
    is_complete = TRUE
WHERE id = $1
  AND is_complete = FALSE
RETURNING *;

-- name: DeleteFile :exec
DELETE
FROM files
WHERE id = $1;

-- name: AttachFileToVersion :exec
INSERT INTO versions_files (version_id, file_id)
VALUES ($1, $2);

-- name: DetachFileFromVersion :exec
DELETE
FROM versions_files
WHERE version_id = $1
  AND file_id = $2;
