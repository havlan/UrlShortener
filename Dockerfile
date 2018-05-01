FROM scratch

ADD urlshort urlshort

ENV PORT 8080

EXPOSE 8080

CMD ["/urlshort"]