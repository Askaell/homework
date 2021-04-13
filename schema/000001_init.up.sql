CREATE TABLE items
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    price numeric not null
);