FROM golang:latest

RUN apt update -y
RUN apt install -y build-essential sudo

RUN go get -u -d gocv.io/x/gocv
RUN cd $GOPATH/src/gocv.io/x/gocv
RUN make install

WORKDIR /go/src/app
COPY . .

CMD ["sh"]

