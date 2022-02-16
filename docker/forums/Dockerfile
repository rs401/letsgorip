FROM scratch

COPY ./.env ./
COPY ./forumsvc ./app/

EXPOSE ${FORUMSVC_PORT}

ENTRYPOINT ["/app/forumsvc"]