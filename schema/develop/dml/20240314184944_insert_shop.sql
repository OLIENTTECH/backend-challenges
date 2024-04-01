-- +goose Up
-- +goose StatementBegin
INSERT INTO shops (
  id,
  name,
  created_at,
  updated_at,
  deleted_at
) VALUES 
    ('01F9ZG3ZZW8Y3VW0KR1H7ZE84T', 'ショップ名1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01F9ZG3XJ90TPTKBK9FJGHK4QY', 'ショップ名2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01F9ZG3TQM2X7VMP8Z9M7P0TZ2', 'ショップ名3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
