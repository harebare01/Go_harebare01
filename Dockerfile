FROM golang:1.7-onbuild
#ENV  http_proxy "http://proxy.xxxx.co.jp:8080/"
#ENV  https_proxy "http://proxy.xxxx.co.jp:8080/"

ONBUILD COPY . /go/src/app
ENTRYPOINT ["go", "run", "main.go"]
#CMD go run main.go
