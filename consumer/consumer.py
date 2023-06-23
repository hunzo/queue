import pika
import json
import os
from pprint import pprint
from dotenv import load_dotenv


import requests
import names

load_dotenv()

QUEUE_NAME = os.environ.get("QUEUE_NAME", "loadtest")
PUBLISH_SERVER = os.environ.get("PUBLISH_AMQP_HOST", "10.10.31.188")
USERNAME = os.environ.get("AMQP_USER", "guest")
PASSWORD = os.environ.get("AMQP_PASSWORD", "guest")
HOSTNAME = os.environ.get("HOSTNAME", "guest")
HOST_LOGGER = os.environ.get("HOST_LOGGER", "guest")

connection_parameter = pika.ConnectionParameters(
    PUBLISH_SERVER,
    heartbeat=600,
    blocked_connection_timeout=300,
    credentials=pika.PlainCredentials(
        USERNAME,
        PASSWORD
    )
)

connection = pika.BlockingConnection(connection_parameter)
channel = connection.channel()
channel.queue_declare(queue=QUEUE_NAME, durable=True)


def Logger(payload):
    url = os.environ.get("HOST_LOGGER", "http://localhost:12201/gelf")
    print(url)
    return requests.post(url, json=payload)


def callback(ch, method, properties, body):
    print(f"Received in  queue_name: {QUEUE_NAME}")

    data = json.loads(body)

    data["host"] = HOSTNAME
    data["name"] = names.get_full_name()

    ret = Logger(data)

    print(ret.status_code)

    pprint(data)


channel.basic_consume(
    queue=QUEUE_NAME, on_message_callback=callback, auto_ack=True)

pprint("Started Consuming...")
pprint({
    "queue_name": QUEUE_NAME,
    "server": PUBLISH_SERVER,
    "username": USERNAME,
    "password": PASSWORD,
    "host_logger": HOST_LOGGER
})

channel.start_consuming()
