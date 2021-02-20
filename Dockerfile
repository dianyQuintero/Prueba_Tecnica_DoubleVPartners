FROM golang:1.14.6-alpine3.12 as builder
COPY go.mod go.sum /go/src/github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners/
WORKDIR /go/src/github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners
RUN go mod download
COPY . /go/src/github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/Prueba_Tecnica_DoubleVPartners github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/dianyQuintero/Prueba_Tecnica_DoubleVPartners/build/Prueba_Tecnica_DoubleVPartners /usr/bin/Prueba_Tecnica_DoubleVPartners
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/Prueba_Tecnica_DoubleVPartners"]