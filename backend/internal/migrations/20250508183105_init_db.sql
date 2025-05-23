-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA texts_schema;

CREATE TABLE texts_schema.user
(
    user_id              BIGSERIAL PRIMARY KEY,
    login                TEXT UNIQUE    NOT NULL,
    password TEXT NOT NULL,
    role TEXT UNIQUE NOT NULL,
    deleted BOOLEAN DEFAULT false
);


CREATE TABLE texts_schema.tag
(
    tag_id   BIGSERIAL PRIMARY KEY,
    tag_name  TEXT UNIQUE      NOT NULL,
    deleted BOOLEAN DEFAULT false
);

CREATE TABLE texts_schema.post
(
    post_id   BIGSERIAL PRIMARY KEY,
    name  TEXT      NOT NULL,
    author TEXT NOT NULL,
    genre TEXT NOT NULL,
    content TEXT NOT NULL,
    editor_id    BIGINT,
    tag_id BIGINT NOT NULL,
    status TEXT DEFAULT 'on_check',
    deleted BOOLEAN DEFAULT false,
    FOREIGN KEY (editor_id) REFERENCES texts_schema.user(user_id),
    FOREIGN KEY (tag_id) REFERENCES texts_schema.tag(tag_id)
);

CREATE TABLE texts_schema.commentary
(
    commentary_id   BIGSERIAL PRIMARY KEY,
    user_id  BIGINT      NOT NULL,
    commentary_content TEXT NOT NULL,
    deleted BOOLEAN DEFAULT false,
    FOREIGN KEY (user_id) REFERENCES texts_schema.user(user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS texts_schema.commentary;
DROP TABLE IF EXISTS texts_schema.post;
DROP TABLE IF EXISTS texts_schema.tag;
DROP TABLE IF EXISTS texts_schema.user;
DROP SCHEMA IF EXISTS texts_schema;
-- +goose StatementEnd