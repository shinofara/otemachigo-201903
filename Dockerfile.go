FROM golang:1.12

WORKDIR /work
ADD main.go main.go

ENTRYPOINT ["go"]
CMD ["run", "main.go"]
