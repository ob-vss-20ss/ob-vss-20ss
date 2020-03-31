FROM obraun/golang-protoactor-ci as builder
COPY . /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor
WORKDIR /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor
RUN go build -o server/server server/server.go

FROM alpine 
COPY --from=builder /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/server/server /app/server
EXPOSE 8091
ENTRYPOINT ["/app/server"]
CMD ["--bind=server:8091"]