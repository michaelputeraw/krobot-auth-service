-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id varchar(36) UNIQUE,
    full_name varchar(255),
    gender varchar(255),
    email varchar(255),
    password varchar(255),
    created_at BIGINT DEFAULT UNIX_TIMESTAMP() NOT NULL,
    updated_at BIGINT DEFAULT UNIX_TIMESTAMP() NOT NULL,
    deleted_at BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users IF EXISTS;
-- +goose StatementEnd
