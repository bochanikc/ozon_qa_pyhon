from flask import Flask, request, render_template, redirect

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
        print(Comment(author, comment))
        return redirect("/comments")

    return render_template("index.html")


@app.route("/comments", methods=["GET"])
def comments():
    example = [Comment('asvezh', 'random text')]
    return render_template("comments.html", comments=example)


if __name__ == '__main__':
    app.run(host="0.0.0.0", debug=True)
