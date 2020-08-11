FROM golang:1.14.6
WORKDIR /usr/src/app
COPY . .
RUN go get "github.com/mattn/go-sqlite3"
RUN go get "github.com/satori/go.uuid"

EXPOSE 3030
CMD ["make","all"]
