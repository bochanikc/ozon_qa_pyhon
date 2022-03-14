import psycopg2
from psycopg2 import Error

DB_URL = 'postgresql://user1:qwerty123@localhost:5432/postgres'


def connect():
    connection = psycopg2.connect(user="user1",
                                  password="qwerty123",
                                  host="localhost",
                                  port="5432",
                                  database="postgres")

    cursor = connection.cursor()
    return connection, cursor


def create_test_table():
    try:
        connection, cursor = connect()
        # SQL query to create a new table
        create_table_query = '''CREATE TABLE animals
                                (
                                    id serial not null primary key,
                                    name varchar(40),
                                    age  integer,
                                    country varchar(40),
                                    place integer
                                );'''
        # Execute a command: this creates a new table
        cursor.execute(create_table_query)
        connection.commit()
        print("Table created successfully")
        insert_table_query = '''INSERT INTO animals (name, age, country, place)
                                VALUES ('Lion', 10, 'Johannesburg', 32),
                               ('Makaka', 3, 'Cape Town', 15),
                               ('Elephant', 32, 'Johannesburg', 2),
                               ('Spider', 2, 'South Town', 3),
                               ('Python', 3, 'North Town', 32),
                               ('Lion 2', 3, 'Johannesburg', 31),
                               ('Lion 3', 2, 'Johannesburg', 32),
                               ('Lion 4', 5, 'Johannesburg', 32),
                               ('Lion 5', 3, 'South Town', 32);'''
        cursor.execute(insert_table_query)
        connection.commit()
        print("Add data to table")

    except (Exception, Error) as error:
        print("Error while connecting to PostgreSQL", error)
        drop_test_table()
    finally:

        if connection:
            cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")


def select_test_table():
    try:
        connection, cursor = connect()
        # SQL query to create a new table
        test_table_query = '''SELECT name, age, country, place FROM animals
                                WHERE name = 'Elephant';'''
        cursor.execute(test_table_query)
        record = cursor.fetchall()
        print("Result ", record)

    except (Exception, Error) as error:
        print("Error while connecting to PostgreSQL", error)
        drop_test_table()
        return None
    finally:
        if connection:
            cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")
            return record


def drop_test_table():
    try:
        connection, cursor = connect()
        drop_table_query = 'drop table animals;'
        cursor.execute(drop_table_query)
        connection.commit()
        print("Drop table successfully")

    except (Exception, Error) as error:
        print("Error while connecting to PostgreSQL", error)
    finally:
        if connection:
            cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")


def test_sql_table():
    create_test_table()
    record = select_test_table()
    if record is not None:
        assert record[0][0] == 'Elephant'
        assert record[0][1] == 32
        assert record[0][2] == 'Johannesburg'
        assert record[0][3] == 2
    drop_test_table()
