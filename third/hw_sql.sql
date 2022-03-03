CREATE DATABASE database;


CREATE TABLE users
(
    id   serial,
    name varchar(40),
    age  integer
);

INSERT INTO users
VALUES (3,'Mariya',23);

------------ 1. Самый взрослый пользователь -----------
SELECT name FROM users
WHERE age = (SELECT max(age) FROM users);



-------------------------------
CREATE TABLE users1
(
    id          serial not null primary key,
    name        varchar(40),
    location_id integer not null references location (id)
);
CREATE TABLE location
(
    id          serial primary key,
    name        varchar(40)
);

INSERT INTO users1 (name, location_id)
VALUES ('Mariya', 1),
       ('Pavel2', 2),
       ('Pavel2', 2),
       ('Pavel', 3),
       ('Pavel1', 1),
       ('Pavel', 2),
       ('Pavel1', 1),
       ('Pavel', 2);

INSERT INTO location (name)
VALUES ('Moscow'),
        ('Stp'),
       ('Borrow');

------------ 2. 10 имен пользователей, которые живут в Москве -----------
SELECT * FROM users1
JOIN location l on l.id = users1.location_id
WHERE l.name = 'Moscow'
limit 10;