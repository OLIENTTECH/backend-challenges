-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "stocks" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "item_id" char(26) NOT NULL,
    "shop_id" char(26) NOT NULL,
    "quantity" INT NOT NULL,
    "expirated_at" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        item_id
    ) REFERENCES items (id) ON DELETE CASCADE,
    FOREIGN KEY (
        shop_id
    ) REFERENCES shops (id) ON DELETE CASCADE
);
CREATE INDEX idx_stocks_created_at ON stocks (created_at);
CREATE INDEX idx_stocks_expirated_at ON stocks (expirated_at);


COMMENT ON COLUMN stocks.id IS '在庫ID';
COMMENT ON COLUMN stocks.item_id IS '商品マスタID';
COMMENT ON COLUMN stocks.shop_id IS '店舗ID';
COMMENT ON COLUMN stocks.quantity IS '残個数';
COMMENT ON COLUMN stocks.expirated_at IS '消費期限';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
