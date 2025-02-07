FROM golang:1.23-bullseye AS build

WORKDIR /app

COPY backend ./
RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app

FROM debian:bullseye-slim

COPY --from=build /bin/app /bin
COPY backend/.env /.env

EXPOSE 8080

CMD [ "/bin/app" ]