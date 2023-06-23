from loads.decorator import excution_time
from datetime import datetime
from uuid import uuid4
from dotenv import load_dotenv

import requests
import pika
import json

import os

import time

load_dotenv()


HOST = os.environ.get("HOST", "localhost")
USERNAME = os.environ.get("USERNAME", "localhost")
PASSWORD = os.environ.get("PASSWORD", "localhost")
QUEUE_NAME = os.environ.get("QUEUE_NAME", "localhost")
HOSTNAME = os.environ.get("HOSTNAME", "publisher-server")


# @excution_time
def tasks(num):
    start_time = time.time()
    uid = str(uuid4())

    print(f"task: {str(uid)}")
    time_now = datetime.now().strftime("%d/%m/%y %H:%M:%S")
    for i in range(num):
        # print(i)
        Logger(i, uid, str(time_now))

    take_time = time.time() - start_time

    print(f"time take: {take_time}")


# @excution_time
def tasks_queue(num):
    start_time = time.time()
    connect = pika.BlockingConnection(
        pika.ConnectionParameters(
            heartbeat=600,
            host=HOST,
            credentials=pika.PlainCredentials(
                USERNAME,
                PASSWORD
            )
        )
    )

    channel = connect.channel()
    channel.queue_declare(
        queue=QUEUE_NAME,
        durable=True
    )

    uid = str(uuid4())
    time_now = datetime.now().strftime("%d/%m/%y %H:%M:%S")
    print(f"tasks_queue: {str(uid)}")
    print(time_now)

    for i in range(num):

        payload = {
            "version": "1.1",
            "host": HOSTNAME,
            "short_message": uid,
            # "full_message": "full message",
            "start_time": str(time_now),
            "count":  i
        }
        channel.basic_publish(
            exchange="",
            routing_key=QUEUE_NAME,
            body=json.dumps(payload),
            properties=pika.BasicProperties(
                content_type="application/json",
                delivery_mode=pika.DeliveryMode.Transient
            )
        )

    take_time = time.time() - start_time

    print(take_time)

    return "success"


def Logger(count, uid, start_time):

    url = os.environ.get("HOST_LOGGER", "http://localhost:12201/glef")
    payload = {
        "version": "1.1",
        "host": "host-loadtest",
        "short_message": uid,
        "full_message": "full message",
        "start_time": start_time,
        "count":  count
    }

    return requests.post(url, json=payload)
