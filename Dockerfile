# Stage 1: Build Golang application
FROM golang:1.20-alpine AS golang-builder

WORKDIR /app

COPY ./backend .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Build Node.js application
FROM node:19-alpine AS node-builder

# Define the build arguments
ARG GTAG
ARG LOCATION

# Now assign the ARG values to environment variables
ENV REACT_APP_GA_TRACKING_ID=${GTAG}
ENV REACT_APP_LOCATION=${LOCATION}

WORKDIR /app

COPY ./frontend .

RUN npm install
RUN npm run build node-sass

# Final stage: Ubuntu base image
FROM ubuntu:20.04

WORKDIR /app

# Copy Golang binary from builder stage
COPY --from=golang-builder /app/main .

# Copy Node.js build from builder stage
COPY --from=node-builder /app/build ./public
COPY --from=node-builder /app/package.json ./public

# Install node and npm in ubuntu
RUN apt-get update && \
    apt-get install -y curl

EXPOSE 8080

CMD ./main