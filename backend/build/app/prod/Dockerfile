FROM golang:1.21.2-alpine3.18

ENV PROJECT_ROOT=/usr/env/
WORKDIR ${PROJECT_ROOT}

COPY ../../../go.mod ${PROJECT_ROOT}
COPY ../../../go.sum ${PROJECT_ROOT}
RUN go mod download

COPY ../../ ${PROJECT_ROOT}

RUN go build -o /go-grpc-todo

CMD [ "/go-grpc-todo" ]
