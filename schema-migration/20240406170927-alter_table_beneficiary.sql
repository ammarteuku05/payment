
-- +migrate Up
ALTER TABLE beneficiaries
RENAME COLUMN beneficiary_bank TO beneficiary_bank_id;

ALTER TABLE beneficiaries
ADD CONSTRAINT fk_beneficiary_bank_id
FOREIGN KEY (beneficiary_bank_id)
REFERENCES beneficiary_banks (id);
-- +migrate Down
