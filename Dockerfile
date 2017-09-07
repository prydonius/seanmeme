FROM busybox

COPY static/ /static/
COPY seanmeme /

ENV PORT 8080

CMD ["/seanmeme"]
