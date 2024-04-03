-- +goose Up
-- +goose StatementBegin
INSERT INTO sales (
    id, 
    barcode_id,
    finance_id, 
    purchased_at,
    created_at, 
    updated_at, 
    deleted_at
) VALUES
    ('01HTDX81GY4KJRNH9G4DN4RW3A', '01HTDPT93Q31JJ7NCDTXV06HRZ', '01FC3W5N4KZ3T9R0D2F9V2MJGX', '2023-11-01 16:01:48.208576', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01HTDX81GY8EXAZMEYNWQYWD2H', '01HTDPT93Q7Y54T3K5P5SP1J6C', '01FC3W5N4P6F8XMPBZQP2BFWGB', '2023-09-22 07:28:08.208576', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01HTDX81GY7Z1YPQ4A2FE214BC', '01HTDPT93QQ8ECW8ZV04P42XPW', '01FC3W5N4RZF7TKH8H10V4J2HY', '2024-03-27 03:00:44.208576', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
