FROM golang:1.19.3-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV dbhost=
ENV dbport=
ENV dbname=
ENV dbuser=
ENV dbpass=
ENV rconip=
ENV rconport=
ENV rconpass=
ENV webip=
ENV webport=

RUN go build -o gorcon app/*.go
CMD ["sh", "-c", "/app/gorcon -dbhost=${dbhost} -dbport=${dbport} -dbname=${dbname} -dbuser=${dbuser} -dbpass=${dbpass} -rconip=${rconip} -rconport=${rconport} -rconpass=${rconpass} -webip=${webip} -webport=${webport}"]