# build stage
FROM golang:alpine AS build-stage
ADD . /src
RUN cd /src/mirrorFinder && go build -o mirrorFinder

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-stage /src/mirrorFinder/mirrorFinder /app/
ENTRYPOINT ./mirrorFinder