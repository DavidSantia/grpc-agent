FROM scratch

ENV NEW_RELIC_LICENSE_KEY=""
EXPOSE 9000

ADD server /
ADD rootca.tar /

ENTRYPOINT ["/server"]
