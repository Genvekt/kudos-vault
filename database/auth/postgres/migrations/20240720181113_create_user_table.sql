-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users"
(
    id            text      NOT NULL PRIMARY KEY,
    first_name    text      NOT NULL,
    last_name     text      NOT NULL,
    email         text      NOT NULL,
    role          text      NOT NULL,
    status        text      NOT NULL,
    password_hash text      NOT NULL,
    created_at    timestamp NOT NULL DEFAULT now(),
    updated_at    timestamp NOT NULL DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
