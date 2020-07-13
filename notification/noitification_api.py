#!/usr/bin/python3
#pylint: disable = invalid-name, broad-except
"""
A notification application which runs on scheduled basis and send the information to users.
Author:- Opstree Solutions
"""

import argparse
import os
import sys
import logging
import time
import emails
import config_with_yaml as config
from elasticsearch import Elasticsearch
import schedule

CONFIG_FILE = os.environ.get("CONFIG_FILE")
FORMATTER = logging.Formatter("%(asctime)s — %(name)s — %(levelname)s — %(message)s")

def init_logger():
    """Function to initialize logger"""
    console_handler = logging.StreamHandler(sys.stdout)
    console_handler.setFormatter(FORMATTER)
    return console_handler

def get_logger():
    """Function to get logger"""
    logger = logging.getLogger("notification-service")
    logger.setLevel(logging.DEBUG)
    logger.addHandler(init_logger())
    return logger

def read_configuration():
    """This function will read the complete configuration file"""
    logger = get_logger()
    try:
        cfg = config.load(CONFIG_FILE)
        return cfg
    except Exception as e:
        logger.error("Not able to parse the configuration file: %s", e)

def send_mail(email_id):
    """function which will send mail to user"""
    logger = get_logger()
    config_content = read_configuration()
    try:
        message = emails.html(
            html="<strong>Your salary slip is generated please check</strong>",
            subject="Salary Slip",
            mail_from=config_content.getProperty("smtp.from"),
        )

        message.send(
            to=email_id,
            smtp={
                "host": config_content.getProperty("smtp.smtp_server"),
                "port": config_content.getProperty("smtp.smtp_port"),
                "timeout": 5,
                "user": config_content.getProperty("smtp.username"),
                "password": config_content.getProperty("smtp.password"),
                "tls": True,
            },
        )
    except Exception as e:
        logger.error("Not able to send the mail: %s", e)

def send_mail_to_all_users():
    """This function will fetch user information from elasticsearch"""
    logger = get_logger()
    config_content = read_configuration()

    try:
        es_client = Elasticsearch(
            [config_content.getProperty("elasticsearch.host")],
            http_auth=(config_content.getProperty("elasticsearch.username"), \
                config_content.getProperty("elasticsearch.password")),
            scheme="http",
            port=config_content.getProperty("elasticsearch.port"),
        )

        result = es_client.search(
            index="employee-management",
            body={
                "query": {
                    "match_all": {}
                }
            }
        )

        for data in result["hits"]["hits"]:
            send_mail(data["_source"]["email_id"])

    except Exception as e:
        logger.error("Error while executing elasticsearch query: %s", e)

def schedule_operation():
    """This function will gets executed for a scheduled interval"""
    logger = get_logger()
    schedule.every().hour.do(send_mail_to_all_users)
    while True:
        logger.info("Waiting for scheduled time to come, I am sure it will come")
        schedule.run_pending()
        time.sleep(1)

if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument("-m", "--mode", \
        help="Mode in which application will run, options - scheduled and external", \
        default="scheduled")
    args = parser.parse_args()

    if args.mode == "scheduled":
        schedule_operation()
    else:
        send_mail_to_all_users()
