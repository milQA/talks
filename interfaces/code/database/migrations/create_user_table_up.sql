create table users
(
    id      bigserial    not null,
    name    varchar(255) not null,
    email   varchar(255) not null,
    address jsonb
);