FROM obraun/golang-protoactor-ci as builder
COPY . /go/src/github.com/ob-vss-20ss/ob-vss-20ss/proto.actor
WORKDIR /go/src/github.com/ob-vss-20ss/ob-vss-20ss/proto.actor
RUN go build -o client/client client/client.go

FROM alpine
COPY --from=builder /go/src/github.com/ob-vss-20ss/ob-vss-20ss/proto.actor/client/client /app/client
EXPOSE 8090
ENTRYPOINT ["/app/client"]
CMD ["--bind=client:8090", "--remote=server:8091"]