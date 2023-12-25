DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS files;

CREATE TABLE users
(
    id          SERIAL,
    email       VARCHAR(50) NOT NULL UNIQUE,
    secret      TEXT        NOT NULL UNIQUE,
    created_at  TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    accessed_at TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

CREATE INDEX users_email_idx ON users (email);

CREATE TABLE files
(
    id          TEXT,
    filename    TEXT      NOT NULL,
    salt        TEXT,
    user_id     INTEGER   NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    accessed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);

ALTER TABLE files
    ADD CONSTRAINT files_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE;
