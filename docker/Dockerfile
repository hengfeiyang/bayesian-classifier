FROM docker.io/golang:1.5.1

RUN go get -u github.com/huichen/sego
RUN go get -u github.com/safeie/bayesian-classifier

COPY docker-entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

EXPOSE 8812
