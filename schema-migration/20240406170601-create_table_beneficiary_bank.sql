
-- +migrate Up
CREATE TABLE beneficiary_banks
(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    beneficiary_bank_name VARCHAR(255) NOT NULL,
    created_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at           TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6)
);

INSERT INTO beneficiary_banks (id, beneficiary_bank_name,created_at,updated_at)
VALUES (
    '12538030-4e59-11ed-8b9d-0242ac120005',
    'Bank Republik Indonesia',
    DEFAULT, -- Automatically set the current timestamp
    DEFAULT  -- Automatically set the current timestamp
),
VALUES (
    '12538030-4e59-11ed-8b9d-0242ac120014',
    'Bank Nasional Indonesia',
    DEFAULT, -- Automatically set the current timestamp
    DEFAULT  -- Automatically set the current timestamp
);

-- +migrate Down
