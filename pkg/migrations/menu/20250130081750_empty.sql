-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA menu;
CREATE SCHEMA IF NOT EXISTS migrations;
CREATE TYPE menu.category AS ENUM ('snacks', 'salads','soups','hot_dishes','side_dishes');
CREATE TABLE menu.chef (
                                 name text
);

CREATE TABLE menu.dishes (
                                id serial primary key,
                                name text not null,
                                category  menu.category not null,
                                favorite bool default false
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
