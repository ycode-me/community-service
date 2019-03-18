FROM golang:1.12.1

WORKDIR $GOPATH/src/community-service

COPY . .

RUN go get ./...

RUN go install

EXPOSE 50051

CMD ["community-service"] 
