FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY handler/*.go ./handler/
COPY shortener/*.go ./shortener/
COPY store/*.go ./store/
COPY pkg ./pkg

RUN go mod tidy
RUN go build -o /urlshortener
RUN apk add tree
RUN tree

EXPOSE 4000

CMD [ "/urlshortener" ]