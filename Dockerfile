# Multistage Build

### CREATE DOCKERMASTER USER
FROM alpine:3.6 AS alpine
RUN adduser -D -u 10001 dockmaster

## MAIN IMAGE
FROM scratch
LABEL Name=products-api
LABEL Author=davyj0nes

COPY --from=alpine /etc/passwd /etc/passwd

ADD products-api /
USER dockmaster

EXPOSE 8080
CMD ["./products-api"]
