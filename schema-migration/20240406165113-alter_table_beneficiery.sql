
-- +migrate Up
ALTER TABLE beneficiaries
ADD COLUMN beneficiary_amount double precision NOT NULL DEFAULT 0;
-- +migrate Down
