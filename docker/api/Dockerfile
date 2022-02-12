FROM scratch

# Need to add ENV variables with sane defaults or no default

COPY ./.env ./
# COPY ./*.pem ./cert/
COPY ./api ./app/

EXPOSE ${API_PORT}

ENTRYPOINT ["/app/api"]