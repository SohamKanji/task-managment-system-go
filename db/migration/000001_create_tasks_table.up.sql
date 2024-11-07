CREATE TABLE "tasks" (
  "id" BIGSERIAL PRIMARY KEY,
  "due" date NOT NULL,
  "status" varchar NOT NULL,
  "priority" bigint NOT NULL,
  "title" varchar NOT NULL,
  "description" varchar NOT NULL
);
