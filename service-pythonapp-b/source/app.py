from flask import Flask, request
import datetime

app = Flask(__name__)


@app.route("/api", methods=["GET"])
def GETAPIDATA():
    return "message from service-b"


@app.route("/api/data", methods=["POST"])
def APIDATA():
    datetime_now = datetime.datetime.now()
    print(str(request.data))
    print(datetime_now.strftime("%Y-%m-%d %H:%M:%S"))
    return "ok"
