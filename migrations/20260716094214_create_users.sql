-- +goose Up
CREATE TABLE users (
    id BIGSERIAL NOT NULL,
    first_mame varchar(255) NOT NULL,
    last_mame varchar(255),
    family_mame varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE users;