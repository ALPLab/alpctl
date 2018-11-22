FROM golang:1.11.4 AS builder
RUN mkdir -p /ALPLab/
WORKDIR /ALPLab/

COPY . .

RUN go mod download && go install ./...
RUN GOOS=windows GOARCH=386 go install ./...
RUN GOOS=darwin GOARCH=amd64 go install ./...

FROM debian:jessie-slim
RUN mkdir -p /release/win && mkdir -p /release/linux  && mkdir -p /release/mac

COPY --from=builder /go/bin/alp /release/linux/
COPY --from=builder /go/bin/windows_386/alp.exe /release/win/
COPY --from=builder /go/bin/darwin_amd64/alp /release/mac/

ENTRYPOINT ["/release/linux/alp"]
