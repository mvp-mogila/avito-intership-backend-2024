-- +goose Up
-- +goose StatementBegin

-- banner table
CREATE TABLE banner (
	id serial PRIMARY KEY,
	content jsonb DEFAULT NULL,
    is_active bool NOT NULL,
	created_at timestamp DEFAULT now() NOT NULL,
	updated_at timestamp DEFAULT now() NOT NULL
);

-- function for banner updated_at
CREATE OR REPLACE FUNCTION update_time()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- trigger for banner updated_at 
CREATE TRIGGER updator
BEFORE UPDATE ON banner
FOR EACH ROW EXECUTE FUNCTION update_time();

--banner_definition table
CREATE TABLE banner_definition (
	banner_id int NOT NULL REFERENCES banner(id) ON DELETE CASCADE ON UPDATE CASCADE,
	feature_id int NULL,
	tag_id int NULL,
	CONSTRAINT banner_definition_pk PRIMARY KEY (feature_id,tag_id)
);

-- banner_id index
CREATE INDEX banner_id_idx ON banner_definition(banner_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER if EXISTS updator ON banner;
DROP FUNCTION IF EXISTS update_time;
DROP INDEX IF EXISTS banner_id_idx;
DROP TABLE IF EXISTS banner_definition;
DROP TABLE IF EXISTS banner;

-- +goose StatementEnd
