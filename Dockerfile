FROM golang

WORKDIR /go/src/app

RUN go get -u github.com/kardianos/govendor

COPY . /go/src/app
#RUN go get -v
RUN bash ./get.sh
RUN go install -v

CMD /go/bin/app