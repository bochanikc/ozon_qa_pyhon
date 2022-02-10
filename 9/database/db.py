import records

DB_URL = 'postgresql://user1:qwerty123@localhost:5432/postgres'


def connect(db_url):
    db = records.Database(db_url=db_url)
    return db


def query_db(db, query):
    rows = db.query(query).as_dict()
    return rows


db = connect(DB_URL)
query = 'CREATE TABLE animals(id serial not null primary key, name varchar(40));'
query_db(db, query)
query = "INSERT INTO animals (name)VALUES ('Moscow'),('Stp'),('Borrow');"
data = query_db(db, query)
print(data)

