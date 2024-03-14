-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "items" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "price" INT NOT NULL,
    "category_id" char(26) NOT NULL,
    "maker_id" char(26) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        category_id
    ) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (
        maker_id
    ) REFERENCES makers (id) ON DELETE CASCADE,
    UNIQUE (name, maker_id)
);
CREATE INDEX idx_items_created_at ON items (created_at);

COMMENT ON COLUMN items.id IS '商品マスタID';
COMMENT ON COLUMN items.name IS '商品名';
COMMENT ON COLUMN items.price IS '単価';
COMMENT ON COLUMN items.category_id IS 'カテゴリID';
COMMENT ON COLUMN items.maker_id IS 'メーカーID';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
