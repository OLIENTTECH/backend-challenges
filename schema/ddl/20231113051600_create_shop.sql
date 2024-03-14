-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "shops" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL
);
CREATE INDEX idx_shops_created_at ON shops (created_at);

COMMENT ON COLUMN shops.id IS '店舗ID';
COMMENT ON COLUMN shops.name IS '店舗名';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS shops;
-- +goose StatementEnd
