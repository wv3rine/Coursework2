-- +goose Up
-- +goose StatementBegin
ALTER TABLE texts_schema.commentary
ADD COLUMN post_id BIGINT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE texts_schema.commentary
DROP COLUMN post_id;
-- +goose StatementEnd
