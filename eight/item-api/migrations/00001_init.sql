-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE items
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(40) NOT NULL,
    description TEXT        NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS items;
