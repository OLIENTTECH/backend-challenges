-- +goose Up
-- +goose StatementBegin
INSERT INTO finances (
    id, 
    shop_id, 
    total_amount, 
    purchased_at, 
    created_at, 
    updated_at, 
    deleted_at
) VALUES 
    ('01FC3W5N4KZ3T9R0D2F9V2MJGX', '01F9ZG3ZZW8Y3VW0KR1H7ZE84T', 5000000, '2021-07-15T12:00:00Z', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01FC3W5N4P6F8XMPBZQP2BFWGB', '01F9ZG3XJ90TPTKBK9FJGHK4QY', 7500000, '2021-07-16T15:30:00Z', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01FC3W5N4RZF7TKH8H10V4J2HY', '01F9ZG3TQM2X7VMP8Z9M7P0TZ2', 3000000, '2021-07-17T09:45:00Z', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
