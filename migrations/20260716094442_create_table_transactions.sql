-- +goose Up
CREATE TYPE TRANSACTION_TYPE as enum ('income', 'expense');

CREATE TABLE transactions (
    id BIGSERIAL NOT NULL,
    code varchar(255) NOT NULL,
    user_id int  NOT NULL,
    category_id int  NOT NULL,
    type TRANSACTION_TYPE NOT NULL,
    price money NOT NULL DEFAULT 0,
    date timestamp NOT NULL,
    PRIMARY KEY(id),

    CONSTRAINT transaction_user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id),

    CONSTRAINT transaction_category_id
        FOREIGN KEY(category_id)
            REFERENCES categories(id)
);

-- +goose Down
DROP TABLE transactions;