ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor

WORKDIR /src

COPY ./ ./

RUN go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/

COPY --from=builder /app /app

EXPOSE 8080

USER nobody:nobody

ENTRYPOINT ["/app"]