FROM centos:latest

COPY ./registry-meta /bin/registry-meta

EXPOSE 9820

ENTRYPOINT ["/bin/registry-meta"]
