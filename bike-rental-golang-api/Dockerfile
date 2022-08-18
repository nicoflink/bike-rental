ARG USER=bikeapp
ARG USER_ID=1000

# Stage BASE: Get Dependencies #
FROM golang:1.16-alpine AS base

# Run apk update
RUN apk update

# Install gcc (necessary for race condition check later on)
RUN apk add build-base

# Install git - to fetch go dependencies
RUN apk add git

# Install ca-certificates - to use https
RUN apk add --no-cache ca-certificates

# Install tzdata for timezones
RUN apk add tzdata


# Disable gcc by default
ENV CGO_ENABLED=0

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

# Stage QA
FROM base AS qa

RUN go test ./... -cover -count 1
RUN CGO_ENABLED=1 go build -race cmd/bike-rental-server/main.go

# Stage BUILD
FROM base AS builder
ARG USER_ID
ARG USER

# create group
RUN addgroup -g ${USER_ID} -S ${USER}

# create user
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${USER_ID}" "${USER}" \
    --ingroup "${USER}"

# create app dir
RUN mkdir app

RUN  go build -o /app/bike-rental-api cmd/bike-rental-server/main.go

# set permissions in app to r-x for dedicated user only
RUN chmod -R 500 /app

# user own app
RUN chown -R ${USER_ID}:${USER_ID} /app

#Stage FINAL
FROM scratch AS final
ARG USER_ID

# import trusted root CAs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# import the user and group files from the build-base
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# by default a scratch image is not provided with zoneinfo for timezones
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# copy application into target directory
COPY --from=builder /app/bike-rental-api /app/bike-rental-api

USER ${USER_ID}

CMD [ "/app/bike-rental-api" ]

EXPOSE 8080