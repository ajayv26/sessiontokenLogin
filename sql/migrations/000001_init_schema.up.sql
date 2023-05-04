BEGIN;

-- Trigger
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- USERS
CREATE TABLE "users" (
    "id" bigserial UNIQUE NOT NULL PRIMARY KEY,
    "code" varchar UNIQUE NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone" varchar UNIQUE NOT NULL,
    "password_hash" varchar NOT NULL,
    "is_admin" boolean NOT NULL DEFAULT FALSE,
    "is_archived" boolean NOT NULL DEFAULT FALSE,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE "otp_sessions" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "user_id" bigint NOT NULL REFERENCES users (id),
    "token" varchar NOT NULL,
    "is_valid" boolean NOT NULL DEFAULT FALSE,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON otp_sessions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

CREATE TABLE "auth_sessions" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "user_id" bigint NOT NULL REFERENCES users (id),
    "token" uuid UNIQUE NOT NULL,
    "is_valid" boolean NOT NULL DEFAULT FALSE,
    "expires_at" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW(),
    "updated_at" timestamptz NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON auth_sessions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();


COMMIT;