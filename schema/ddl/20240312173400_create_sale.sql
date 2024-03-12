-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sales (
    id char(26) NOT NULL PRIMARY KEY,
    barcode_id char(26) NOT NULL,
    account_id char(26) NOT NULL,
    purchase_date TIMESTAMP NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (
        barcode_id
    ) REFERENCES itemBarcodes (id) ON DELETE CASCADE,
    FOREIGN KEY (
        account_id
    ) REFERENCES accounts (id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sales;
-- +goose StatementEnd
