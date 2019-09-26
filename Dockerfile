FROM alpine
ADD web-test  /web-test
WORKDIR /
ENTRYPOINT [ "/web-test" ]
