FROM golang:latest

LABEL maintainer="Kiriakos Papachristou <kiriakos.papachristou@gmail.com>"

# Build Arguments
ARG APP_NAME=welcome-demo-app
#ARG LOG_DIR=/${APP_NAME}/logs

# Create Log Direcotry
#RUN mkdir -p ${LOG_DIR}

# Environment variables
#ENV LOG_FILE_LOCATION=${LOG_DIR}/app.log

# Set the Current Working Directory inside the container
WORKDIR /go/src/welcome-demo-app
# Copy everything from the current directory to pwd inside the container
COPY . . 
# Download dependencies
RUN go get -d -v ./...
# Install the package
RUN go install -v ./...

# Expose on this port the application
EXPOSE 6969

# Declare additional volumes to mount
#VOLUME ["welcome-demo-app/logs"]

# Run the binary file produced by `go install`
CMD ["welcome-demo-app"]