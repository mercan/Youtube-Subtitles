FROM golang:1.18-alpine

RUN apk upgrade
RUN apk add --no-cache curl
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -o /usr/local/bin/yt-dlp
RUN chmod a+rx /usr/local/bin/yt-dlp

# install dependencies
RUN apk update \ && apk add python3-dev
RUN ln -s /usr/bin/python3 /usr/local/bin/python

# Set the Current Working Directory inside the container
WORKDIR /usr/app

# Copy the go.mod and go.sum files
COPY go.* ./

# Download all the dependencies
RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . ./

# Build the Go app
RUN go build -o ./GoYoutubeSubtitles

# This container exposes port 3000 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./GoYoutubeSubtitles"]