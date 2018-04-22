FROM golang

COPY . /go/src/app
WORKDIR /go/src/app
RUN go get -v
RUN go install -v

CMD /go/bin/app