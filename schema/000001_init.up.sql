CREATE TABLE users
(
    id            serial       not null unique,
    guid          INT          DEFAULT NULL,
    ip            varchar(255) DEFAULT NULL,
    token         varchar(255) DEFAULT NULL,
    time          INT          DEFAULT NULL
);