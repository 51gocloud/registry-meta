FROM scratch

COPY ./registry-meta /bin/registry-meta

EXPOSE 6000

ENTRYPOINT ["/bin/registry-meta"]
