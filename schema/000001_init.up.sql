CREATE TABLE item
(
    id serial not null unique,
    name varchar(255) not null,
    description varchar(255) not null,
    price numeric not null,
    discountPrice numeric not null,
    discount real,
    dayItem boolean,
    vendorCode varchar (255) not null,
    category varchar (255) not null,
    amount varchar(255) not null
);