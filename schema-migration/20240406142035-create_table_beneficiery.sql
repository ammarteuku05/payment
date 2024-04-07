
-- +migrate Up
CREATE TABLE beneficiaries
(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    beneficiary_name VARCHAR(255) NOT NULL,
    beneficiary_account_number VARCHAR(255) NOT NULL,
    beneficiary_bank VARCHAR(255) NOT NULL,
    created_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);
-- +migrate Down
