-- +goose Up
-- +goose StatementBegin
INSERT INTO users (
  id,
  login_id,
  password,
  family_name,
  given_name,
  role_id
) VALUES (
  '11edf3a8-2264-d984-bd6f-0242ac120003',
  '0123456789',
  'P@ssw0rd',
  'logisco',
  'tarou',
  1
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
