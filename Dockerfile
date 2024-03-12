FROM golang:latest

WORKDIR /server

COPY . .

RUN go mod tidy

RUN chmod +x ./scripts/build.sh

RUN ./scripts/build.sh

CMD ["./build/server"]