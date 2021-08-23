FROM mysql:8.0.26
LABEL maintainer="kevyn24ec@gmail.com"
ENV MYSQL_ROOT_PASSWORD @36a0dd17
COPY ./database/tables /home/mysql/tvdb/Tables
COPY ./database/scripts /home/mysql/tvdb/Scripts
WORKDIR /home/mysql/tvdb/
COPY ./database/setup.sql /docker-entrypoint-initdb.d/