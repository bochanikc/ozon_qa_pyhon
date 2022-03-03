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


def sql_table():
    try:
        connection, cursor = connect()
        # SQL query to create a new table
        create_table_query = '''CREATE TABLE animals
        (id serial not null primary key, name varchar(40));'''
        # Execute a command: this creates a new table
        cursor.execute(create_table_query)
        connection.commit()
        print("Table created successfully in PostgreSQL ")

    except (Exception, Error) as error:
        print("Error while connecting to PostgreSQL", error)
    finally:
        if connection:
            cursor.close()
            connection.close()
            print("PostgreSQL connection is closed")


# conn = connect(DB_URL)
sql_table()

# query_db("CREATE TABLE animals(id serial not null primary key, name varchar(40));")
# print(rows)
# rows = query_db("INSERT INTO animals (name)VALUES ('Moscow'),('Stp'),('Borrow');")
# print(rows)
