FROM postgres

COPY migrations/000001_create_tables.up.sql /docker-entrypoint-initdb.d/create.sql
COPY migrations/000002_seeds.up.sql /docker-entrypoint-initdb.d/seed.sql