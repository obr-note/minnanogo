FROM python:3.6-alpine

RUN apk update && apk add strace
