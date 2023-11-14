FROM gcr.io/distroless/static:nonroot

COPY uptimectl /usr/local/bin/uptimectl
ENTRYPOINT ["/usr/local/bin/uptimectl"]
