-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "categories" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL
);
CREATE INDEX idx_categories_created_at ON categories (created_at);

COMMENT ON COLUMN categories.id IS 'カテゴリID';
COMMENT ON COLUMN categories.name IS 'カテゴリ名';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
