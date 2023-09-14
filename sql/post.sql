CREATE TABLE Post (
    Id uuid PRIMARY KEY,
    UserId uuid NOT NULL,
    
    Title text NOT NULL,
    "description" text NOT NULL,
    Images text[] NOT NULL,

    "status" boolean NOT NULL,
    Price int NOT NULL,
    Tags   text NOT NULL,
    
    "time" timestamp NOT NULL
);


CREATE TABLE View (
    Id int PRIMARY KEY,
    Count int,
    PostId uuid unique references Post(Id) ON DELETE CASCADE
);

CREATE TABLE View_To_User (
    Id int PRIMARY KEY,
    ViewId uuid references View(Id) ON DELETE CASCADE,
    UserId uuid unique references Post(Id) ON DELETE SET NULL
);
