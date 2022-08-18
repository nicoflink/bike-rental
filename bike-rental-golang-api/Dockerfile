# Stage BASE: Get Dependencies #
FROM golang:1.16-alpine AS base

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bike-rental-api cmd/bike-rental-server/main.go

CMD [ "./bike-rental-api" ]

EXPOSE 8080

# Stage QA
FROM base AS qa

RUN go test ./... -cover -count 1
RUN go go build -race cmd/bike-rental-server/main.go



# Stage BUILD
FROM base AS builder

RUN go build -o /bike-rental-api cmd/bike-rental-server/main.go

#Stage FINAL
FROM scratch AS final

COPY --from=build /bike-rental-api /bike-rental-api

CMD [ "./bike-rental-api" ]

EXPOSE 8080