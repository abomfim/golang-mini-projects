
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id      CHAR(36)      NOT NULL,
    firstname    VARCHAR(255)                NOT NULL,
    lastname VARCHAR(255)                NOT NULL,
    username VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE users;