FROM scratch

COPY ./.env ./
COPY ./placesvc ./app/

EXPOSE ${PLACESVC_PORT}

ENTRYPOINT ["/app/placesvc"]