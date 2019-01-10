FROM python:3.7-alpine

VOLUME /tmp

ARG PROJECT_NAME=app

COPY . /app

ENTRYPOINT ["gunicorn" "${PROJECT_NAME}.wsgi:application" "--bind" "0.0.0.0:8000"]rake db:migrate