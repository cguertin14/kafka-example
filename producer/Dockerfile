# Step 1 - compile code binary
FROM golang:1.20.6-alpine AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -o ./program .


# Step 2 - import necessary files to run program.
FROM gcr.io/distroless/base-debian11:nonroot
COPY --from=builder /app/program /program
ENTRYPOINT ["/program"]