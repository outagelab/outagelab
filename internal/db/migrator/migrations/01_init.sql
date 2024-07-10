-- +goose Up

CREATE TABLE IF NOT EXISTS "accounts"
(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "data" JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS "api_keys"
(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "data" JSONB NOT NULL
);

CREATE TABLE IF NOT EXISTS "users"
(
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "email" VARCHAR NOT NULL,
    "data" JSONB NOT NULL
);

CREATE INDEX ON "users" ("email");

-- +goose Down

DROP TABLE "accounts";
DROP TABLE "api_keys";
DROP TABLE "users";
