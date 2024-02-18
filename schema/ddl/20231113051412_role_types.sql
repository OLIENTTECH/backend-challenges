-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_types (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_types;
-- +goose StatementEnd
