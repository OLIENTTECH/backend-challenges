-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS itemBarcodes (
    id char(26) NOT NULL PRIMARY KEY,
    item_id char(26) NOT NULL,
    regisration_date TIMESTAMP NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (
        item_id
    ) REFERENCES items (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS itemBarcodes;
-- +goose StatementEnd
