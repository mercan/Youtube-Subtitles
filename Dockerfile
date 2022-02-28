# syntax=docker/dockerfile:1
FROM golang:1.17-alpine

RUN apk upgrade
RUN apk add --no-cache curl
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl
RUN chmod a+rx /usr/local/bin/youtube-dl

# install dependencies
RUN apk update \ && apk add python3-dev
RUN ln -s /usr/bin/python3 /usr/local/bin/python

WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY ../asdad .

# Download all the dependencies
RUN go mod download

# Build the application
RUN go build -o /Go-Youtube-Subtitles

EXPOSE 3000

# Run the application
CMD [ "/Go-Youtube-Subtitles" ]