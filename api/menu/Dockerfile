FROM golang:latest 
EXPOSE 3000
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
ENV GOPATH /app
RUN go get github.com/codegangsta/negroni
RUN go get github.com/gocql/gocql
RUN go get github.com/gorilla/mux
RUN go get github.com/unrolled/render
RUN go build main.go 
CMD ["./main"]

