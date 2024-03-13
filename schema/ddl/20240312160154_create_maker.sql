-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS makers (
    "id" char(26) NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz DEFAULT NULL
);
CREATE INDEX idx_makers_created_at ON makers (created_at);

COMMENT ON COLUMN makers.id IS 'メーカーID';
COMMENT ON COLUMN makers.name IS 'メーカー名';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS makers;
-- +goose StatementEnd
