CREATE TABLE "event" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "date_time" timestamp NOT NULL,
  "venue" VARCHAR(255),
  "type" VARCHAR(255),
  "organizer_id" INTEGER
);

CREATE TABLE "organizer" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "company" VARCHAR(255),
  "logo" BYTEA,
  "address" TEXT
);

CREATE TABLE "attendee" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "event_id" INTEGER,
  "ticket_number" VARCHAR(255),
  "payment_status" BOOLEAN,
  "payment_date" DATE
);

CREATE TABLE "ticket" (
  "id" SERIAL PRIMARY KEY,
  "event_id" INTEGER,
  "attendee_id" INTEGER,
  "type" VARCHAR(255),
  "price" "NUMERIC(10, 2)" NOT NULL
);

CREATE TABLE "speaker" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "event_id" INTEGER,
  "profile_image" BYTEA,
  "bio" TEXT
);

CREATE TABLE "session" (
  "id" SERIAL PRIMARY KEY,
  "title" VARCHAR(255) NOT NULL,
  "description" TEXT,
  "date_time" timestamp NOT NULL,
  "venue" VARCHAR(255),
  "event_id" INTEGER,
  "speaker_id" INTEGER
);

CREATE TABLE "sponsorship" (
  "id" SERIAL PRIMARY KEY,
  "event_id" INTEGER,
  "company_name" VARCHAR(255) NOT NULL,
  "contact_person" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "package" VARCHAR(255)
);

CREATE TABLE "exhibition" (
  "id" SERIAL PRIMARY KEY,
  "event_id" INTEGER,
  "company_name" VARCHAR(255) NOT NULL,
  "contact_person" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "booth_number" VARCHAR(255),
  "floor_plan" BYTEA
);

CREATE TABLE "volunteer" (
  "id" SERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "email" VARCHAR(255) NOT NULL,
  "phone" VARCHAR(255),
  "event_id" INTEGER,
  "role" VARCHAR(255),
  "availability" TEXT
);

CREATE INDEX ON "event" ("date_time");

CREATE INDEX ON "event" ("organizer_id");

CREATE INDEX ON "organizer" ("company");

CREATE INDEX ON "attendee" ("event_id");

CREATE INDEX ON "attendee" ("payment_status");

CREATE INDEX ON "attendee" ("payment_date");

CREATE INDEX ON "ticket" ("event_id");

CREATE INDEX ON "speaker" ("event_id");

CREATE INDEX ON "session" ("event_id");

CREATE INDEX ON "session" ("speaker_id");

CREATE INDEX ON "session" ("date_time");

CREATE INDEX ON "sponsorship" ("event_id");

CREATE INDEX ON "sponsorship" ("package");

CREATE INDEX ON "exhibition" ("event_id");

CREATE INDEX ON "volunteer" ("event_id");

CREATE INDEX ON "volunteer" ("role");

ALTER TABLE "event" ADD FOREIGN KEY ("organizer_id") REFERENCES "organizer" ("id");

ALTER TABLE "attendee" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "ticket" ADD FOREIGN KEY ("attendee_id") REFERENCES "attendee" ("id");

ALTER TABLE "speaker" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("speaker_id") REFERENCES "speaker" ("id");

ALTER TABLE "sponsorship" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "exhibition" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");

ALTER TABLE "volunteer" ADD FOREIGN KEY ("event_id") REFERENCES "event" ("id");
