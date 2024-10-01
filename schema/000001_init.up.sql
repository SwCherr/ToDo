CREATE TABLE users
(
    "id"            SERIAL        PRIMARY KEY,
    "email"         varchar(255)  NOT NULL,
    "password"      varchar(255)  NOT NULL
);

CREATE TABLE refresh_sessions (
    "id"            SERIAL                      PRIMARY KEY,
    "user_id"       INT                         NOT NULL,
    "ip"            varchar(15)                 NOT NULL,
    "guid"          varchar(255) UNIQUE         NOT NULL,
    "refresh_token" varchar(255) UNIQUE         NOT NULL,
    "expires_in"    INTEGER                     NOT NULL
);
