-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS match_id_seq

-- Table Definition
CREATE TABLE "public"."matchs" (
    "id" int4 NOT NULL DEFAULT nextval('match_id_seq'::regclass),
    "sport_key" varchar,
    "sport_nice" varchar,
    "commence_time" int4,
    "home_team" varchar,
    "team_a" varchar,
    "team_b" varchar,
    "sites_count" int4
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS odds_id_seq

-- Table Definition
CREATE TABLE "public"."odds" (
    "id" int4 NOT NULL DEFAULT nextval('odds_id_seq'::regclass),
    "match_key" varchar NOT NULL,
    "site_key" varchar NOT NULL,
    "site_nice" varchar NOT NULL,
    "last_update" int4 NOT NULL,
    "odds_team_a" float4 NOT NULL,
    "odds_team_b" float4 NOT NULL,
    "odds_draw" float4 NOT NULL
);

-- This script only contains the table creation statements and does not fully represent the table in database. It's still missing: indices, triggers. Do not use it as backup.

-- Squences
CREATE SEQUENCE IF NOT EXISTS sport_id_seq

-- Table Definition
CREATE TABLE "public"."sports" (
    "id" int4 NOT NULL DEFAULT nextval('sport_id_seq'::regclass),
    "sport_key" varchar NOT NULL,
    "active" bool NOT NULL,
    "sport_group" varchar NOT NULL,
    "detail" varchar NOT NULL,
    "title" varchar NOT NULL
);






