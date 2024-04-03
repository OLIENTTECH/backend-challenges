-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "sales" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "barcode_id" char(26) NOT NULL,
    "finance_id" char(26) NOT NULL,
    "purchased_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        barcode_id
    ) REFERENCES item_barcodes (id) ON DELETE CASCADE,
    FOREIGN KEY (
        finance_id
    ) REFERENCES finances (id) ON DELETE CASCADE
);
CREATE INDEX idx_sales_created_at ON sales (created_at);

COMMENT ON COLUMN sales.id IS '売上ID';
COMMENT ON COLUMN sales.barcode_id IS 'バーコードID';
COMMENT ON COLUMN sales.finance_id IS '会計ID';
COMMENT ON COLUMN sales.purchased_at IS '購入日';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sales;
-- +goose StatementEnd
