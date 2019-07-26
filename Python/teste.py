import flask
from flask import Flask, render_template

app = Flask(__name__, template_folder = '')

@app.route("/")
def index():
    return "Flask App!"

@app.route("/hello/<string:name>/")
def hello(name):
    return render_template('hello.html',name = name)

if __name__ == "__main__":
    app.run(host = '0.0.0.0', port = 3000)