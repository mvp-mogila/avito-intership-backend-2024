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

--banner_definition table
CREATE TABLE banner_definition (
	banner_id int NOT NULL REFERENCES banner(id) ON DELETE CASCADE ON UPDATE CASCADE,
	feature_id int NULL,
	tag_id int NULL,
	CONSTRAINT banner_definition_pk PRIMARY KEY (feature_id,tag_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS banner_definition;
DROP TABLE IF EXISTS banner;

-- +goose StatementEnd
