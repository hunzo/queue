FROM python:3.10-alpine as builder

COPY . /app
WORKDIR /app

RUN apk update \
    && apk add --no-cache

RUN python3 -m venv /opt/venv
ENV Path="/opt/venv/bin:$PATH"

COPY requirements.txt .
RUN /opt/venv/bin/pip install --no-cache --upgrade pip && \
    /opt/venv/bin/pip install -r requirements.txt && \
    chmod -x /app/entrypoint.sh

CMD ["entrypoint.sh"]