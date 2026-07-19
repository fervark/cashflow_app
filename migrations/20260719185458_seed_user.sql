-- +goose Up
SELECT 'up SQL query';
INSERT INTO users (first_name, last_name, family_name)
VALUES ('John', 'Doe', 'Too');

-- +goose Down
SELECT 'down SQL query';
DELETE FROM users ['first_name' = 'John', 'last_name' = 'Doe', 'family_name' = 'Too']