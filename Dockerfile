FROM golang:1.11.1

WORKDIR /go/src/alp
COPY . .
ADD VERSION .

RUN go get -d -v github.com/spf13/cobra \
github.com/spf13/viper \
github.com/gogo/protobuf/proto \
google.golang.org/grpc \
golang.org/x/net/context \
github.com/inconshreveable/mousetrap

RUN go install -v github.com/spf13/cobra \
github.com/spf13/viper \
github.com/gogo/protobuf/proto \
google.golang.org/grpc \
golang.org/x/net/context \
github.com/inconshreveable/mousetrap

CMD ["alp"]
