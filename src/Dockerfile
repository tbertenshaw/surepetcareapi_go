ARG GO_VERSION=1.17
 
# STAGE 1: building the executable
FROM golang:${GO_VERSION}-alpine AS build
WORKDIR /build
COPY *.go /build/
COPY ./go.mod /build/

RUN CGO_ENABLED=0 go build -o /petsafeapi ./


FROM scratch as final
#FROM alpine:latest as final
COPY --from=build /petsafeapi /petsafeapi

# copy ca certs
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 9090
CMD ["/petsafeapi"]