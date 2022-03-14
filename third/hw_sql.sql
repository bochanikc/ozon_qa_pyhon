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


----------- 3. Подготовка данных для теста на Python -------------------
CREATE TABLE animals
(
    id serial not null primary key,
    name varchar(40),
    age  integer,
    country varchar(40),
    place integer
);

INSERT INTO animals (name, age, country, place)
VALUES ('Lion', 10, 'Johannesburg', 32),
       ('Makaka', 3, 'Cape Town', 15),
       ('Elephant', 32, 'Johannesburg', 2),
       ('Spider', 2, 'South Town', 3),
       ('Python', 3, 'North Town', 32),
       ('Lion 2', 3, 'Johannesburg', 31),
       ('Lion 3', 2, 'Johannesburg', 32),
       ('Lion 4', 5, 'Johannesburg', 32),
       ('Lion 5', 3, 'South Town', 32);


SELECT name, age, country, place FROM animals
WHERE name = 'Elephant';