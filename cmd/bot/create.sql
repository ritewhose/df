DROP TABLE IF EXISTS Quote;
DROP TABLE IF EXISTS User;
DROP TABLE IF EXISTS Tell;

CREATE TABLE Tell (
    targetUser  TEXT,
    tellMsg     TEXT,

    FOREIGN KEY(targetUser) REFERENCES User(userName)
);

CREATE TABLE Quote (
    creator     TEXT,
    quoteMsg    TEXT,

    FOREIGN KEY(creator) REFERENCES User(userName)
);

CREATE TABLE User (
    userName    TEXT PRIMARY KEY
);
