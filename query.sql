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

-- name: ListFiles :many
SELECT files.id,
       files.created_at,
       files.updated_at,
       files.name,
       files.size,
       files.path,
       files.mime_type,
       files.is_complete
FROM files
         LEFT JOIN versions_files ON files.id = versions_files.file_id
WHERE (sqlc.narg('versionId')::BIGINT IS NULL OR versions_files.version_id = sqlc.narg('versionId'))
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
INSERT INTO files (name, size, path, mime_type, is_complete)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateFile :one
UPDATE files
SET updated_at  = current_timestamp,
    name        = $2,
    size        = $3,
    path        = $4,
    mime_type   = $5,
    is_complete = $6
WHERE id = $1
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

-- Users
-- name: GetUserById :one
SELECT id,
       created_at,
       updated_at,
       name,
       email,
       email_verified
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByProvider :one
SELECT users.id,
       users.created_at,
       users.updated_at,
       users.name,
       users.email,
       users.email_verified
FROM users
    JOIN public.external_auth ea on users.id = ea.user_id
WHERE ea.provider = $1 AND ea.provider_id = $2
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, email, email_verified)
VALUES ($1, $2, $3)
RETURNING *;

-- ExternalAuths
-- name: ListExternalAuthsByUserId :many
SELECT
    id,
    created_at,
    updated_at,
    user_id,
    provider,
    provider_id
FROM external_auth
WHERE user_id = $1;

-- name: CreateExternalAuth :one
INSERT INTO external_auth (user_id, provider, provider_id)
VALUES ($1, $2, $3)
RETURNING *;
