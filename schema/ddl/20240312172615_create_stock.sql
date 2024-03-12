-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stocks (
    id char(26) NOT NULL PRIMARY KEY,
    item_id char(26) NOT NULL,
    store_id char(26) NOT NULL,
    quantity INT NOT NULL,
    expiration_date TIMESTAMP,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (
        item_id
    ) REFERENCES items (id) ON DELETE CASCADE,
    FOREIGN KEY (
        store_id
    ) REFERENCES stores (id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stocks;
-- +goose StatementEnd
