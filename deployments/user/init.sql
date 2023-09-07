CREATE TABLE users (
    Id uuid PRIMARY KEY,
    Email text NOT NULL UNIQUE,
    PhoneNumber text NOT NULL UNIQUE,
    Login text NOT NULL UNIQUE,
    Password text NOT NULL,
    Name text NOT NULL,
    PathToAvatar text
);