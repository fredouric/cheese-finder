-- +goose Up
CREATE TABLE IF NOT EXISTS cheese (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    departement VARCHAR(255) NOT NULL,
    fromage varchar(255) not null,
    pagefrancaise VARCHAR(255) NOT NULL,
    englishpage VARCHAR(255) NOT NULL,
    lait VARCHAR(255) NOT NULL,
    geoshape  VARCHAR(255) NOT NULL,
    geopoint2d VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE cheese;
