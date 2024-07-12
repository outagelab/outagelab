# backend
FROM golang:1.22 AS backend

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o dist/outagelab app.go

# frontend
FROM node:20-alpine3.19 AS frontend
WORKDIR /app
COPY ui/package.json ui/yarn.lock ./
RUN yarn
COPY ui/ ./
RUN yarn build-only --outDir dist

# runtime
FROM alpine

WORKDIR /app

COPY --from=backend /app/dist/outagelab /app/outagelab
COPY --from=frontend /app/dist/ /app/dist

EXPOSE 8080

CMD ["/app/outagelab"]
