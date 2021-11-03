##
## Build
##

#TODO change base image
FROM  golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY *.go ./
RUN go build -o /docker-gs-ping

##
## Deploy
##

FROM golang:1.16-alpine
WORKDIR /
COPY --from=build /docker-gs-ping /docker-gs-ping
EXPOSE 8080
ENTRYPOINT ["/docker-gs-ping"]
