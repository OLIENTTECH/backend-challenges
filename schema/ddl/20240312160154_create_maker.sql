-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS makers (
    id char(26) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS makers;
-- +goose StatementEnd
