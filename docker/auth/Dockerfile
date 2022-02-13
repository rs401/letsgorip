FROM scratch

# COPY ./.env ./
COPY ./authsvc ./app/

EXPOSE ${AUTHSVC_PORT}

ENTRYPOINT ["/app/authsvc"]