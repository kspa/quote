################################
# STEP 1 build executable binary
################################
FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR $GOPATH/src/github.com/kspa/quote
COPY . .

# Fetch dependencies.
# Using go get.
#RUN go get -d -v

# Using go mod.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/quote

############################
# STEP 2 build a small image
############################
FROM scratch

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/quote /go/bin/quote

# Use an unprivileged user.
USER appuser:appuser

# Port on which the service will be exposed.
EXPOSE 8080

# Run the quote binary.
ENTRYPOINT ["/go/bin/quote"]