-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "item_barcodes" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "item_id" char(26) NOT NULL,
    "registrated_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        item_id
    ) REFERENCES items (id) ON DELETE CASCADE
);
CREATE INDEX idx_item_barcodes_created_at ON item_barcodes (created_at);

COMMENT ON COLUMN item_barcodes.item_id IS '商品マスタID';
COMMENT ON COLUMN item_barcodes.registrated_at IS '登録日';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_barcodes;
-- +goose StatementEnd
