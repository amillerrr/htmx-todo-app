FROM golang:1.21.4-bookworm AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./db/*.go ./db/
COPY ./handlers/*.go ./handlers/
COPY ./models/*.go ./models/
COPY ./services/*.go ./services/
COPY ./utils/*.go ./utils/

RUN go build -o /htmx-todo-app

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=build /htmx-todo-app /htmx-todo-app
COPY ./templates/index.html ./templates/

USER nonroot:nonroot

ENTRYPOINT ["/htmx-todo-app"]