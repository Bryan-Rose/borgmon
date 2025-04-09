FROM node:lts-bookworm AS buildfrontend
WORKDIR /build
COPY frontend/package.json \
    frontend/package-lock.json \
    frontend/svelte.config.js \
    frontend/vite.config.ts \
    frontend/tsconfig.json ./

RUN npm install

COPY frontend/. .
RUN npm run build .


FROM golang:1.24-bookworm AS buildbackend
WORKDIR /build

RUN curl -o borg -L --silent https://github.com/borgbackup/borg/releases/download/1.4.0/borg-linux-glibc236 \
    && chmod +x ./borg

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
COPY --from=buildfrontend build/ frontend/build/
RUN go build -o borgmon backend/main.go && chmod +x borgmon


FROM debian:bookworm AS final
WORKDIR /app
EXPOSE 8090

COPY --from=buildbackend /build/borgmon /build/borg ./

VOLUME /Borgmon-data
ENTRYPOINT [ "/app/borgmon" ]
CMD [ "serve" ]