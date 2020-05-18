FROM golang:1.7-onbuild
#ONBUILD COPY . /go/src/app
ENTRYPOINT ["go", "run", "main.go"]

