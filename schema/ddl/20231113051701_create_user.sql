-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id char(36) PRIMARY KEY,
    login_id VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    family_name VARCHAR(255) NOT NULL,
    given_name VARCHAR(255) NOT NULL,
    role_id INT NOT NULL,
    last_logined_at TIMESTAMP DEFAULT NULL,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (
        role_id
    ) REFERENCES role_types (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
