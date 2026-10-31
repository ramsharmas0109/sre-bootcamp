FROM golang:1.26-alpine

WORKDIR /student-api

COPY . .

RUN apk add --no-cache make curl; \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.20.1/migrate.linux-amd64.tar.gz | tar xvz; \
    mv migrate /usr/bin; \
    go mod tidy; \
    make build \

#RUN make migrate_up


#requirement- 

#  1. machine with go installed. - done
#  2. code, package, env - done
#  3. golang migrate tool.
#  4. make installed. - done
#  5. we need curl to install migrate tool.