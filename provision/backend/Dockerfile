# docker-ce needs to be 17.05 or later.
FROM golang:alpine as go_builder
RUN echo $GOPATH and $GOROOT
COPY vendor/  $GOPATH/src/
RUN ls -lah  $GOPATH/src/
COPY server/  /server
WORKDIR /server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trackerd .

FROM scratch
WORKDIR /server
COPY --from=go_builder /server/trackerd .
COPY provision/backend/config.json .
COPY migrations /migrations/
ENTRYPOINT ["/server/trackerd"]
