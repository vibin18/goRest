# build stage
FROM golang:alpine AS build-stage
ADD . /src
RUN cd /src/api && go build -o api

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-stage /src/api/api /app/
ENTRYPOINT ./mirrorFinder