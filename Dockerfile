FROM golang:latest

# Copy apt.conf for proxy server.
# Remove it if you don't need this.
COPY ./apt.conf /etc/apt/apt.conf

RUN apt update -y
RUN apt install -y build-essential sudo

# The following code is impossible under the proxy environment.
# So you need to execute manually, and then commit its as a new container image.
# Of course you can uncomment if you are not under the proxy environment.
# RUN go get -u -d gocv.io/x/gocv
# RUN cd $GOPATH/src/gocv.io/x/gocv
# RUN make install

WORKDIR /go/src/app
COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...
# CMD ["app"]
CMD ["sh"]

