FROM alpine:latest AS api
RUN apk add --update bash curl && rm -rf /var/cache/apk/*
COPY ./harvest /bin/harvest
EXPOSE 443
EXPOSE 80
CMD ["/bin/harvest"]
