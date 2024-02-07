-- Create "projects" table
CREATE TABLE "projects" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NULL,
  "sequence" numeric NULL,
  "slug" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_projects_deleted_at" to table: "projects"
CREATE INDEX "idx_projects_deleted_at" ON "projects" ("deleted_at");
-- Create "tasks" table
CREATE TABLE "tasks" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "title" text NULL,
  "content" text NULL,
  "start_time" timestamptz NULL,
  "end_time" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_tasks_deleted_at" to table: "tasks"
CREATE INDEX "idx_tasks_deleted_at" ON "tasks" ("deleted_at");
-- Create "project_task" table
CREATE TABLE "project_task" (
  "project_id" bigint NOT NULL,
  "task_id" bigint NOT NULL,
  PRIMARY KEY ("project_id", "task_id"),
  CONSTRAINT "fk_project_task_project" FOREIGN KEY ("project_id") REFERENCES "projects" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_project_task_task" FOREIGN KEY ("task_id") REFERENCES "tasks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
