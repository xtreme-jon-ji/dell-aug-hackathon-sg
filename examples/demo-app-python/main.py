import flask
import os

app = flask.Flask(__name__)

@app.route("/", methods=['GET'])
def index():
    return "Hello World from a Python App!"

if __name__ == "__main__":
  port = int(os.getenv("PORT", 8080))
  app.run(host='0.0.0.0', port=port)