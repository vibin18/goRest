# build stage
FROM golang:alpine AS build-stage
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD .. /src
RUN cd /src && go build -o mirrorFinder

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-stage /src/mirrorFinder /app/
ENTRYPOINT ./mirrorFinder