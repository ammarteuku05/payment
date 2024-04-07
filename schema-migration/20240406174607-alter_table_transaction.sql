
-- +migrate Up
CREATE TABLE transactions
(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    sender_account_number VARCHAR(255) NOT NULL,
    recipient_account_number VARCHAR(255) NOT NULL,
    amount double precision NOT NULL,
    status VARCHAR(20) NULL,
    transaction_date date,
    created_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);
-- +migrate Down
