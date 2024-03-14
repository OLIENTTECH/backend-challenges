-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "users" (
    "id" char(26) NOT NULL PRIMARY KEY,
    "shop_id" char(26) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "is_shop_manager" BOOLEAN NOT NULL,
    "last_logined_at" timestamptz DEFAULT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL,
    FOREIGN KEY (
        shop_id
    ) REFERENCES shops (id) ON DELETE CASCADE
);
CREATE INDEX idx_users_created_at ON users (created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
