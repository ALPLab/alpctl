FROM golang:1.11.4 AS builder
RUN mkdir -p /ALPLab/
WORKDIR /ALPLab/

COPY . .

RUN go mod download && go install ./...
RUN GOOS=windows GOARCH=386 go install ./...

FROM debian:jessie-slim
RUN mkdir -p /release

COPY --from=builder /go/bin/alp /release/
COPY --from=builder /go/bin/windows_386/alp.exe /release/

ENTRYPOINT ["/release/linux/alp"]
