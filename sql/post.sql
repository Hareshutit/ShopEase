CREATE TABLE Post (
    Id uuid PRIMARY KEY,
    UserId uuid NOT NULL,
    "close" boolean NOT NULL,
    Title text NOT NULL,
    "description" text NOT NULL,
    Price text NOT NULL,
    Tags   text NOT NULL,
    Images text[] NOT NULL,
    "time" timestamp NOT NULL,
);


CREATE TABLE Views (
    Count id,
    PostId uuid PRIMARY KEY,
    Id 
);



