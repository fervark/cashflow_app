-- +goose Up
CREATE TABLE categories (
    id BIGSERIAL NOT NULL,
    name varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE categories;