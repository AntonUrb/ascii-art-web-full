FROM golang:1.18.10-bullseye

WORKDIR /app
COPY . .

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "."]

LABEL version="1.0" maintainer="Anton Urban && Daniil Leonov <https://01.kood.tech/git/Anton/ascii-art-web-dockerize>"