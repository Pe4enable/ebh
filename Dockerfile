FROM golang:1.10.2-alpine
ARG github_auth_token
ADD . /go/src/github.com/BankEx/ebh
WORKDIR /go/src/github.com/BankEx/ebh

RUN apk update
RUN apk add git

RUN git config --global url."https://${github_auth_token}:x-oauth-basic@github.com/".insteadOf 'https://github.com/'

RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock

RUN dep ensure --vendor-only

RUN CGO_ENABLED=0 go build -a -o main

FROM alpine
COPY --from=0 /go/src/github.com/BankEx/ebh /
COPY config/config.y*ml /config

EXPOSE 8335 27017
CMD ["/main"]


