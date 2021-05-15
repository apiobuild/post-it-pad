FROM golang:1.14-buster AS build

RUN mkdir /app
ADD . /app
WORKDIR /app

# Running test before build
RUN go test ./... -coverprofile cp.out
# TODO: More compact if build and copy
RUN go build -o pad cmd/pad/main.go

FROM debian:buster-slim
WORKDIR /app
COPY --from=build /app/pad pad

RUN mkdir layouts
COPY layouts layouts

RUN apt-get update

ENTRYPOINT ["./pad"]
CMD ["-h"]