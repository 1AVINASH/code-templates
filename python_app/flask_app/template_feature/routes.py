from flask import request
from flask_restful import Resource

from template_feature.dtos import Input
from utility.logger import app_logger


class Template(Resource):
    def get(self, template_id):
        return {"message ": f"Template with id {template_id} fetched successfully"}

    def create(self):
        payload = Input.CreateTemplate(**request.get_json())
        app_logger.info(f"Received payload for creating template: {payload}")
        return {"message": "Template Created successfully"}

    def update(self):
        payload = Input.UpdateTemplate(**request.get_json())
        app_logger.info(f"Received payload for updating template: {payload}")
        return {"message": "Template Updated successfully"}
