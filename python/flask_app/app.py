import argparse

from flask import Flask, request, Response
from flask_cors import CORS
from flask_restful import Api

from constants.defaults import DEFAULT_HOST, DEFAULT_PORT
from template_feature.routes import Template

app = Flask(__name__)
CORS(app)
api = Api(app)

# Add routes
api.add_resource(Template, "/template", "/template/<int:template_id>", endpoint="template")

if __name__=="__main__":
    parser = argparse.ArgumentParser(description="Start the flask server")
    parser.add_argument("--host", default=DEFAULT_HOST, help=f"Host to bind (default: {DEFAULT_HOST})")
    parser.add_argument("--port", type=int, default=DEFAULT_PORT, help=f"Port to bind (default: {DEFAULT_PORT})")
    args = parser.parse_args()
    app.run(host=args.host, port=args.port, debug=True)