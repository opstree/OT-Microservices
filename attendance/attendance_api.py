#!/usr/local/bin/python
# pylint: disable=no-member, bare-except, line-too-long
"""
Attendance application is part of OT-Microservices.
This application stores and pull attendance data from MySQL
Available Endpoints:-
    - /attendance/create
    - /attendance/search
    - /attendance/healthz
"""

import json
import os
import logging
# import pymysql.cursors
import yaml
from flask import Flask, jsonify, request
from flaskext.mysql import MySQL
from elasticapm.contrib.flask import ElasticAPM
from elasticapm.handlers.logging import LoggingHandler

ELASTIC_APM_SERVER_URL = os.getenv('ELASTIC_APM_SERVER_URL', 'http://localhost:8200')
ELASTIC_APM_SERVICE_NAME = os.getenv('ELASTIC_APM_SERVICE_NAME', 'attendance')
CONFIG_FILE = os.getenv('CONFIG_FILE', '/app/config/config.yaml')

pymysql = MySQL()
app = Flask(__name__)
apm = ElasticAPM(app,
                server_url=ELASTIC_APM_SERVER_URL,
                service_name=ELASTIC_APM_SERVICE_NAME,
                logging=True)

def read_config():
    """Method for reading the configuration file for attendance"""
    with open(CONFIG_FILE, 'r', encoding="utf-8") as config_file:
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

@app.route("/attendance/healthz", methods=['GET'])
def check_health():
    """Method for checking the health of MySQL"""
    try:
        connection = initiate_database()
        connection.ping()
        return jsonify(mysql="up", description="MySQL is healthy"), 200
    except:
        app.logger.error("Unable to make mysql connection", exc_info=True)
        return jsonify(mysql="down", description="MySQL is not healthy"), 400

@app.route("/attendance/create", methods=['POST'])
def push_attendance_data():
    """For pushing attendance data inside MySQL Database"""
    record = json.loads(request.data)
    connection = initiate_database()
    with connection:
        try:
            with connection.cursor() as cursor:
                sql = "CREATE TABLE IF NOT EXISTS Employee ( id int(6) NOT NULL, status varchar(50) NOT NULL, date varchar(50), PRIMARY KEY (id) )"
                cursor.execute(sql)
            connection.commit()
        except:
            app.logger.error("Unable to create table in MySQL", exc_info=True)
            return 400

        try:
            with connection.cursor() as cursor:
                sql = "INSERT INTO `Employee` (`id`, `status`, `date`) VALUES (%s, %s, %s)"
                cursor.execute(sql, (record['id'], record['status'], record['date']))
            connection.commit()
            app.logger.info("Successfully pushed attendance data in MySQL", exc_info=True)
            return jsonify(message="Successfully uploaded the attendance data"), 200
        except:
            app.logger.error("Unable to push attendance data in MySQL", exc_info=True)
            return jsonify(message="Error in pushing attendance data"), 400

@app.route("/attendance/search", methods=['GET'])
def fetch_attendance_data():
    """For pulling attendance data from MySQL Database"""
    connection = initiate_database()
    with connection:
        try:
            with connection.cursor() as cursor:
                sql = "SELECT * FROM Employee ORDER BY id DESC"
                cursor.execute(sql)
                result = cursor.fetchall()
            connection.commit()
            app.logger.info("Successfully pulled attendance data from MySQL", exc_info=True)
            return jsonify(result), 200
        except:
            app.logger.error("Unable to pull attendance data from MySQL", exc_info=True)
            return jsonify(message="Error while pulling data for attendance"), 400


if __name__ == "__main__":
    handler = LoggingHandler(client=apm.client)
    console_handler = logging.StreamHandler()
    app.logger.addHandler(handler)
    app.logger.addHandler(console_handler)
    app.run()
