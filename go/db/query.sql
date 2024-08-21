-- name: AddCheese :exec
INSERT INTO cheese (departement, fromage, pagefrancaise, englishpage, lait, geoshape, geopoint2d) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetCheese :one
SELECT * FROM cheese 
WHERE id=? LIMIT 1;

-- name: GetAllCheeses :many
SELECT * FROM cheese LIMIT ? OFFSET ?;

-- name: DeleteAllCheeses :exec
DELETE FROM cheese;
