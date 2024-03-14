-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "finances" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "shop_id" char(26) NOT NULL,
    "total_amount" INT NOT NULL,
    "purchased_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        shop_id
    ) REFERENCES shops (id) ON DELETE CASCADE
);
CREATE INDEX idx_finances_created_at ON finances (created_at);

COMMENT ON COLUMN finances.id IS '会計ID';
COMMENT ON COLUMN finances.shop_id IS '店舗ID';
COMMENT ON COLUMN finances.total_amount IS '合計金額';
COMMENT ON COLUMN finances.purchased_at IS '購入日';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS finances;
-- +goose StatementEnd
