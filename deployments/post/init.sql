CREATE TABLE posts (
    Id uuid PRIMARY KEY,
    UserId uuid NOT NULL,
    "close" boolean NOT NULL,
    Title text NOT NULL,
    "description" text NOT NULL,
    Price text NOT NULL,
    Tags   text NOT NULL,
    Images text[] NOT NULL,
    "time" timestamp NOT NULL,
    Views integer,
    fts   TSVECTOR
);


CREATE TABLE favorite (
    IdPost uuid[],
    UserId uuid PRIMARY KEY
);

