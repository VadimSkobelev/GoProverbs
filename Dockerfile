FROM golang AS compiling_stage
WORKDIR /go/src/go_proverbs
COPY main.go .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go_proverbs .

FROM alpine:latest
LABEL version="0.1.0"
WORKDIR /root/
COPY --from=compiling_stage /go/src/go_proverbs/go_proverbs .
ENTRYPOINT ./go_proverbs
EXPOSE 27002