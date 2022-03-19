from flask import Flask, request, render_template, redirect
import sqlite3

app = Flask(__name__)


class Comment(object):
    def __init__(self, author, record):
        self.author = author
        self.record = record


@app.route("/", methods=["POST", "GET"])
def index():
    if request.method == "POST":
        author = request.form["author"]
        comment = request.form["comment"]
        insert_comment(author, comment)
        return redirect("/comments")

    return render_template("index.html")


@app.route("/comments", methods=["GET"])
def comments():
    all_results = select_all_comment()
    all_comments = []
    for data in all_results:
        all_comments.append(Comment(data[1], data[2]))
    return render_template("comments.html", comments=all_comments)


def connect_db():
    return sqlite3.connect('data.db')


def create_table():
    conn = connect_db()
    cur = conn.cursor()
    cur.execute(f"""CREATE TABLE IF NOT EXISTS comments(
               id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, author TEXT, comment TEXT);""")
    conn.commit()
    conn.close()


def insert_comment(author, comment):
    conn = connect_db()
    cur = conn.cursor()
    cur.execute(f"""INSERT INTO comments (author, comment) VALUES('{author}', '{comment}');""")
    conn.commit()


def select_all_comment():
    conn = connect_db()
    cur = conn.cursor()
    cur.execute("SELECT * FROM comments;")
    all_results = cur.fetchall()
    print(all_results)
    return all_results


if __name__ == '__main__':
    create_table()
    app.run(host="0.0.0.0", debug=True)
