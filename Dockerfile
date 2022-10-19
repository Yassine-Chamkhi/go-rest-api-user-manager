# Build Stage
FROM golang:1.18.3-alpine as build-env
 
# Set environment variable
ENV APP_NAME go-rest-api-user-manager
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Build application
RUN go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# Run Stage
FROM alpine:latest
 
# Set environment variable
ENV APP_NAME go-rest-api-user-manager
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .
COPY ./migrations/postgres ./migrations/postgres
 
# Expose application port
EXPOSE 8080
 
# Start app
CMD ./$APP_NAME