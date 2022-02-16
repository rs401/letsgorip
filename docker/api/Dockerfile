FROM scratch

COPY ./.env ./
COPY ./api ./app/

EXPOSE ${API_PORT}

ENTRYPOINT ["/app/api"]