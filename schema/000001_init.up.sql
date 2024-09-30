CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    user_ip       varchar(255) DEFAULT NULL,
    refresh_token varchar(255) DEFAULT NULL,
    time_life_rt  varchar(255) DEFAULT NULL
);