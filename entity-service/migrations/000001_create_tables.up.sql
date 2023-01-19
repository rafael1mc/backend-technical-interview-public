CREATE TABLE athletes (
  id SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  league varchar NOT NULL
);

CREATE TABLE teams (
  id SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  league varchar NOT NULL
);

CREATE TABLE nfl_stats (
  id SERIAL PRIMARY KEY,
  entity_id integer NOT NULL,
  entity_type varchar NOT NULL,
  field_goals_made integer NOT NULL DEFAULT 0,
  assists integer NOT NULL DEFAULT 0,
  fumbles integer NOT NULL DEFAULT 0
);

CREATE TABLE nba_stats (
  id SERIAL PRIMARY KEY,
  entity_id integer NOT NULL,
  entity_type varchar NOT NULL,
  three_points_made integer NOT NULL DEFAULT 0,
  two_points_made integer NOT NULL DEFAULT 0,
  free_throws_made integer NOT NULL DEFAULT 0
);
