FROM scratch

# Need to add ENV variables with sane defaults or no default

COPY ./.env ./
# COPY ./*.pem ./cert/
COPY ./authsvc ./app/

EXPOSE ${AUTHSVC_PORT}

ENTRYPOINT ["/app/authsvc"]