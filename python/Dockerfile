FROM python:3.9-alpine

RUN adduser -D worker
USER worker

LABEL language="python"
LABEL org.opencontainers.image.source https://github.com/guionardo/auth-service

WORKDIR /app

COPY --chown=worker:worker ./requirements.txt ./requirements.txt

RUN pip install --no-cache-dir --upgrade -r ./requirements.txt

COPY --chown=worker:worker . .

CMD ["/home/worker/.local/bin/uvicorn", "main:app", "--host", "0.0.0.0", "--port","8000"]
