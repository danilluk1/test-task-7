CREATE TABLE "statistics" (
  "id" serial PRIMARY KEY,
  "user_id" varchar NOT NULL,
  "count" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "statistics" ("id");

CREATE INDEX ON "statistics" ("user_id");
