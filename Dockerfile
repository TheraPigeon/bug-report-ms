FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=auto \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Download dependency and git
RUN apk update && apk add git
RUN go get github.com/joho/godotenv

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 8080

# Initializes web hook env variables
ARG SOUP_HOOK

ENV SOUP_HOOK=$SOUP_HOOK

# Command to run when starting the container
CMD ["/dist/main"]