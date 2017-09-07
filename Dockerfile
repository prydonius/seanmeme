FROM alpine

RUN apk add -U ca-certificates

COPY static/ /static/
COPY seanmeme /

ENV PORT 8080

CMD ["/seanmeme"]
