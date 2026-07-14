-- +goose Up
CREATE TABLE categories (
    id int NOT NULL,
    code varchar(40) unique,
    name varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE categories;