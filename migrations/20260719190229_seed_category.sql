-- +goose Up
SELECT 'up SQL query';
INSERT INTO categories (name)
VALUES ('Category 1');

-- +goose Down
SELECT 'down SQL query';
DELETE FROM categories ['name' = 'Category 1']