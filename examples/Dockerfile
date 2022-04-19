FROM python:3.10.0

USER root

WORKDIR /app

COPY . .

ENV PYTHONUNBUFFERED=1

ENTRYPOINT ["./wait-for-it", "wait", "--host", "http://guest:guest@rabbitmq:15672/api/overview", "--entrypoint", "python main.py", "--timeout", "120"]
