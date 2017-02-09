FROM        sdurrheimer/alpine-glibc:latest
MAINTAINER  Thibault NORMAND <me@zenithar.org>

ADD https://github.com/tianon/gosu/releases/download/1.10/gosu-amd64 /usr/bin/gosu
ADD entrypoint.sh /
ADD bin/server /usr/bin

RUN chmod +x /usr/bin/server \
    && chmod +x /usr/bin/gosu \
    && chmod +x /entrypoint.sh \
    && addgroup nikoniko \
    && adduser -s /bin/false -G nikoniko -S -D nikoniko \
    && mkdir /app \
    && mkdir /data

ADD views /data/views
ADD static /data/static

EXPOSE     5000
WORKDIR    /data
VOLUME     ["/data"]
ENTRYPOINT [ "/entrypoint.sh" ]
CMD        [ "app:help" ]
