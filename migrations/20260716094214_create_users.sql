-- +goose Up
CREATE TABLE users (
    id BIGSERIAL NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255),
    family_name varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;