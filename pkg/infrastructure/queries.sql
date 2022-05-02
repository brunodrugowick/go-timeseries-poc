-- name: GetMeasurement :one
SELECT * FROM measurement
WHERE uuid = $1 LIMIT 1;

-- name: ListMeasurements :many
SELECT * FROM measurement
ORDER BY created_date;

-- name: CreateMeasurement :one
INSERT INTO measurement (
  created_date, heart_rate, high, low, username
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteMeasurement :exec
DELETE FROM measurement
WHERE uuid = $1;

