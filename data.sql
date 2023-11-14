drop table if exists customer cascade;
drop table if exists loader cascade;
drop table if exists task cascade;

create table customer
(
    id serial primary key,
    login varchar(100) not null,
    password varchar(100) not null,
    money int
);

create table task
(
    id serial primary key,
    name varchar(100),
    weight varchar(100) not null,
    done boolean default false
);

create table loader
(
    id serial primary key,
    login varchar(100) not null,
    password varchar(100) not null,
    Weight float,
    money int,
    Drunk  boolean
);

-- http://localhost:8080/register?login=hh&role=customer&password=123