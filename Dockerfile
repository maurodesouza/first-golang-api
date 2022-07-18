FROM golang:1.18 AS build

WORKDIR /app

COPY . .

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]
