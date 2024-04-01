-- +goose Up
-- +goose StatementBegin
INSERT INTO makers (
    id, 
    name, 
    created_at, 
    updated_at, 
    deleted_at
) VALUES 
    ('01F9ZG43E0MV5GJKM5J8D6KY3Z', 'メーカー名1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01F9ZG43M32PRK6JPMKZQ0ZH6X', 'メーカー名2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    ('01F9ZG43JFS85QWEV6P2Q0ZP6W', 'メーカー名3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
