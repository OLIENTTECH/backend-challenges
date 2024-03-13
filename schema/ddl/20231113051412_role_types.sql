-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS role_types (
    "id" char(26) NOT NULL PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP DEFAULT NULL
);
CREATE INDEX idx_role_types_created_at ON role_types (created_at);

COMMENT ON COLUMN role_types.id IS 'ロールID';
COMMENT ON COLUMN role_types.id IS 'ロール名';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS role_types;
-- +goose StatementEnd
