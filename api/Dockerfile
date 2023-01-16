FROM golang:alpine
RUN mkdir /app
ADD ./api/ /app/
WORKDIR /app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./api/go.mod .
COPY ./api/go.sum .
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . 
RUN adduser -S -D -H -h /app appuser 
USER appuser 
EXPOSE 8000 
CMD ["./main"]