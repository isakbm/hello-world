FROM scratch

COPY server/server /server

ENTRYPOINT ["/server"]