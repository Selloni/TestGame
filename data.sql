drop table if exists customer cascade;
drop table if exists loader cascade;
drop table if exists task cascade;
drop table if exists completed_tasks cascade;

create table customer
(
    id serial primary key,
    login varchar(100) not null,
    password varchar(100) not null,
    money integer
);


insert into customer (login, password, money) values ('llc', 123, 902020);

ALTER TABLE customer
    ADD CONSTRAINT customer_constraint UNIQUE (login);

create table task
(
    id serial primary key,
    name varchar(100),
    weight integer not null,
    done boolean default false
);

create table loader
(
    id serial primary key,
    login varchar(100) not null,
    password varchar(100) not null,
    weight integer,
    money integer,
    drunk  boolean,
    tired integer default 0,
    task_id integer ,
    CONSTRAINT task_fk FOREIGN KEY (task_id) REFERENCES completed_tasks (id)
);

ALTER TABLE loader
    ADD CONSTRAINT loader_constraint UNIQUE (login);

CREATE TABLE completed_tasks (
     id serial PRIMARY KEY,
     loader_id integer REFERENCES loader(id),
     task_id integer REFERENCES task(id)
);

-- http://localhost:8080/register?login=hh&role=customer&password=123