# get official base image for the Dockerfile based on Alpine Linux version 1.20
FROM golang:1.20-alpine

# define env variables for the workig project directory path within the container
# enable Go modules with GO111MODULE and disables C compiler support
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# sert the working directory to /app
WORKDIR /app

# creates directory named /build inside the container
RUN mkdir "/build"

# copies all the files from the host to the /app inside the container
COPY . .

# donwloads the go modules inside the go.mod file
RUN go mod download

# Install MongoDB dependencies
RUN apk --no-cache add ca-certificates mongodb-tools

# exposes port to the container
EXPOSE 4444:4444

#  downloads and installs CompileDeamon from github for automatic reload
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon


# execute CompileDeamon command which rebuilds the app usidng 'go build' and outputs the binary to '/build/app' and runs '/build/app'
ENTRYPOINT CompileDaemon -build="go build -o /build/app ./cmd/api" -command="/build/app"