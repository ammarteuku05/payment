
-- +migrate Up
INSERT INTO beneficiaries (id, beneficiary_name,beneficiary_account_number,beneficiary_bank_id,created_at,updated_at)
VALUES (
    '02538030-4e59-11ed-8b9d-0242ac120004',
    'Joh Khannedy',
    '1234567890',
    '12538030-4e59-11ed-8b9d-0242ac120005',
    DEFAULT, -- Automatically set the current timestamp
    DEFAULT  -- Automatically set the current timestamp
);
-- +migrate Down
