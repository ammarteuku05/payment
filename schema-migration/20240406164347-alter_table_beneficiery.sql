
-- +migrate Up
CREATE TYPE beneficieries_status AS ENUM ('active', 'inactive');
ALTER TABLE beneficieries
ADD COLUMN status beneficieries_status NOT NULL DEFAULT 'active';
-- +migrate Down
