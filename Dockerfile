FROM golang:1.19-alpine as development

WORKDIR /usr/src/app

RUN apk add curl

# RUN apk add postgresql

# RUN mkdir /run/postgresql
# RUN su postgres -c 'mkdir /var/lib/postgresql/data'
# RUN su postgres -c 'chmod 0700 /var/lib/postgresql/data'
# RUN su postgres -c 'initdb -D /var/lib/postgresql/data'
# RUN chmod -R 777 /run/postgresql
# RUN su postgres -c 'echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf'

COPY go.mod go.sum ./

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

COPY . .
RUN go mod download && go mod verify

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

ENTRYPOINT ["/entrypoint.sh", "development"]

FROM golang:1.19-alpine as build

WORKDIR /usr/src/app

COPY . .

FROM alpine:3.14 as production

COPY --from=build /usr/src/app/golang-my-app /
COPY ./entrypoint.sh /

RUN apk add dos2unix

RUN chmod +x /golang-my-app

RUN chmod +x /entrypoint.sh

RUN apk add postgresql

RUN mkdir /run/postgresql

RUN su postgres -c 'mkdir /var/lib/postgresql/data'

RUN su postgres -c 'chmod 0700 /var/lib/postgresql/data'

RUN su postgres -c 'initdb -D /var/lib/postgresql/data'

RUN su postgres -c 'echo "host all all 0.0.0.0/0 md5" >> /var/lib/postgresql/data/pg_hba.conf'

COPY postgresql/postgresql.conf /var/lib/postgresql/data/postgresql.conf

RUN chown postgres:postgres /run/postgresql/

EXPOSE 80

ENTRYPOINT ["/entrypoint.sh"]