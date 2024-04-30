-- name: AddCheese :exec
INSERT INTO Cheese (Departement, Fromage, PageFrancaise, EnglishPage, Lait, GeoShape, GeoPoint2D) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetCheese :one
SELECT * FROM Cheese 
WHERE id=? LIMIT 1;
