-- name: GetProject :one
SELECT * FROM projects WHERE id = ? LIMIT 1;

-- name: ListProjects :many
SELECT
    projects.*,
    locations.name AS location_name,
    locations.lat AS location_lat,
    locations.lon AS location_lon
FROM projects
    INNER JOIN locations on locations.id = projects.location_id
ORDER BY projects.created_at;

-- name: ListLocations :many
SELECT * FROM locations ORDER BY created_at DESC;
