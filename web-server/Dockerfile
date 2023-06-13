FROM golang:1.20.3-bullseye

# Set the Working Directory inside the container
WORKDIR /app

# Copy app files
COPY . .
COPY ./static /app/static/

RUN go build -o web-server 

EXPOSE 8080

CMD [ "./web-server" ]

