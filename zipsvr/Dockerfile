FROM alpine
RUN apk add --no-cache ca-certificates
COPY zipsvr /zipsvr
EXPOSE 80
ENTRYPOINT ["/zipsvr"]
