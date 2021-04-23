FROM golang:alpine

WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the code into the container
COPY . /app

# Build the application
RUN go build -o main .

EXPOSE 8008

CMD ["./main"]

# Build docker image:
# docker build -t biblio .

# Run docker image:
# docker run -p 8008:8008 biblio:latest