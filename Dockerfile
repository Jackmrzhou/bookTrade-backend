FROM alpine:latest

RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /goworkspace/bookTrade/
COPY . ./

EXPOSE 8080
ENTRYPOINT ["./bookTrade"]
