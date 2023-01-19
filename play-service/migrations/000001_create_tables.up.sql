CREATE TABLE picks (
  id SERIAL PRIMARY KEY,
  user_id integer NOT NULL,
  entity_id integer NOT NULL,
  entity_type varchar NOT NULL
);