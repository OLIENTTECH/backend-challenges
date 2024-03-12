-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS items (
    id char(26) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    category_id char(26) NOT NULL,
    maker_id char(26) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (
        category_id
    ) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN KEY (
        maker_id
    ) REFERENCES makers (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
