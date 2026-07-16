-- +goose Up
CREATE TYPE TRANSACTION_TYPE as enum ('income', 'expense');

CREATE TABLE transactions (
    id BIGSERIAL NOT NULL,
    code varchar(255) NOT NULL,
    type TRANSACTION_TYPE NOT NULL,
    price money NOT NULL DEFAULT 0,
    date timestamp NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE transactions;