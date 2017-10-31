FROM ubuntu:17.10

WORKDIR /root
ENV GOPATH=/root

RUN apt-get update && apt-get install -y golang-go golang-glide

ADD reflection proto fixtures src/github.com/lazada/grpc-ui/reflection
ADD static src/github.com/lazada/grpc-ui/static
ADD main.go glide.yaml glide.lock src/github.com/lazada/grpc-ui/

RUN cd src/github.com/lazada/grpc-ui && glide install
RUN go install github.com/lazada/grpc-ui

CMD /root/bin/grpc-ui