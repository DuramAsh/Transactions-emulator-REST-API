FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o constanta-golang-emulator-task ./cmd/main.go

CMD ["./constanta-golang-emulator-task"]