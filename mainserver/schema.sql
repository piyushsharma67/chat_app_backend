CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  name varchar      NOT NULL,
  email  varchar NOT NULL,
  password varchar NOT NULL
);