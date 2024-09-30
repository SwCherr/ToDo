CREATE TABLE users
(
    id            serial       not null unique,
    guid          INT          not null unique,
    ip            varchar(255) DEFAULT NULL,
    token         varchar(512) DEFAULT NULL,
    time          INT     DEFAULT NULL
);