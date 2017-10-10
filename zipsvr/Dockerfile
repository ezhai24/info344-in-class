FROM alpine
# RUN apk add --no-cache ca-certificates
COPY zips.csv /zips.csv
COPY zipsvr /zipsvr
EXPOSE 443
ENTRYPOINT ["/zipsvr"]
