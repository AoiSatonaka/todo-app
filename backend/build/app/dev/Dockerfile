FROM golang:1.21.2-alpine3.18

ENV PROJECT_ROOT=/usr/src/
WORKDIR ${PROJECT_ROOT}

RUN go install github.com/cosmtrek/air@latest

COPY ../../../../go.mod ${PROJECT_ROOT}
COPY ../../../../go.sum ${PROJECT_ROOT}
RUN go mod download

CMD ["air", "-c", ".air.toml"]
