-- create "study_groups" table
CREATE TABLE "study_groups" ("id" character varying NOT NULL, "name" character varying NOT NULL, "faculty_name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "teachers" table
CREATE TABLE "teachers" ("id" character varying NOT NULL, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- create "user_configs" table
CREATE TABLE "user_configs" ("id" uuid NOT NULL, "email" character varying NOT NULL, "base" jsonb NOT NULL, "extended_groups" jsonb NOT NULL, "exclude_rules" jsonb NOT NULL, PRIMARY KEY ("id"));
-- create index "userconfig_email" to table: "user_configs"
CREATE UNIQUE INDEX "userconfig_email" ON "user_configs" ("email");
