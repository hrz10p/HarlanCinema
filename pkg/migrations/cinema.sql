CREATE TABLE "seance"(
    "id" BIGINT NOT NULL,
    "movie_id" BIGINT NOT NULL,
    "date" DATE NOT NULL,
    "location" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "seance" ADD PRIMARY KEY("id");
CREATE TABLE "tickets"(
    "user_id" BIGINT NOT NULL,
    "seance_id" BIGINT NOT NULL,
    "cost" BIGINT NOT NULL,
    "ticket_type" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "tickets" ADD PRIMARY KEY("user_id");
ALTER TABLE
    "tickets" ADD PRIMARY KEY("seance_id");
CREATE TABLE "reviews"(
    "user_id" BIGINT NOT NULL,
    "movie_id" BIGINT NOT NULL,
    "text" TEXT NOT NULL,
    "rate" DOUBLE PRECISION NOT NULL
);
ALTER TABLE
    "reviews" ADD PRIMARY KEY("user_id");
ALTER TABLE
    "reviews" ADD PRIMARY KEY("movie_id");
CREATE TABLE "users"(
    "id" BIGINT NOT NULL,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "role" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("id");
CREATE TABLE "movies"(
    "id" BIGINT NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT NOT NULL,
    "rating" DOUBLE PRECISION NOT NULL
);
ALTER TABLE
    "movies" ADD PRIMARY KEY("id");
ALTER TABLE
    "reviews" ADD CONSTRAINT "reviews_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");
ALTER TABLE
    "seance" ADD CONSTRAINT "seance_movie_id_foreign" FOREIGN KEY("movie_id") REFERENCES "movies"("id");
ALTER TABLE
    "tickets" ADD CONSTRAINT "tickets_seance_id_foreign" FOREIGN KEY("seance_id") REFERENCES "seance"("id");
ALTER TABLE
    "reviews" ADD CONSTRAINT "reviews_movie_id_foreign" FOREIGN KEY("movie_id") REFERENCES "movies"("id");
ALTER TABLE
    "tickets" ADD CONSTRAINT "tickets_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");