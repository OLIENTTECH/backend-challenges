-- +goose Up
-- +goose StatementBegin
INSERT INTO role_types (type)
VALUES
(
  'admin'
),
(
  'general'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
