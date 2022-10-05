#!/usr/local/bin/python

import pymysql.cursors
import yaml
import os
from flask import Flask, jsonify
from elasticapm.contrib.flask import ElasticAPM
from elasticapm.handlers.logging import LoggingHandler
import logging

ELASTIC_APM_SERVER_URL = os.getenv('ELASTIC_APM_SERVER_URL', 'http://localhost:8200')
ELASTIC_APM_SERVICE_NAME = os.getenv('ELASTIC_APM_SERVICE_NAME', 'attendance')
CONFIG_FILE = os.getenv('CONFIG_FILE', '/app/config/config.yaml')

app = Flask(__name__)
apm = ElasticAPM(app, server_url=ELASTIC_APM_SERVER_URL, service_name=ELASTIC_APM_SERVICE_NAME, logging=True)

def read_config():
    """Method for reading the configuration file for attendance"""
    with open(CONFIG_FILE, 'r') as config_file:
        yaml_values = yaml.load(config_file, Loader=yaml.FullLoader)
    return yaml_values

def initiate_database():
    """Method for initiating the database connection"""
    config_properties = read_config()
    mysql_connection = pymysql.connect(host=config_properties['mysql']['host'],
                                       user=config_properties['mysql']['username'],
                                       password=config_properties['mysql']['password'],
                                       database=config_properties['mysql']['db_name'],
                                       cursorclass=pymysql.cursors.DictCursor)
    return mysql_connection

@app.route("/health")
def check_health():
    """Method for checking the health of MySQL"""
    try:
        connection = initiate_database()
        connection.ping()
        return jsonify(
            mysql="up",
            description="MySQL is healthy"), 200
    except Exception as e:
        app.logger.error("Unable to make mysql connection", exc_info=True)
        return jsonify(
            mysql="down",
            description="MySQL is not healthy"), 400

if __name__ == "__main__":
    handler = LoggingHandler(client=apm.client)
    console_handler = logging.StreamHandler()
    app.logger.addHandler(handler)
    app.logger.addHandler(console_handler)
    app.run()
