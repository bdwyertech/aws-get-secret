FROM golang:1.15-alpine as aws-get-secret
WORKDIR /go/src/github.com/bdwyertech/aws-get-secret
COPY . .
ARG VCS_REF
RUN CGO_ENABLED=0 GOFLAGS='-mod=vendor' go build -ldflags="-X main.GitCommit=$VCS_REF -X main.ReleaseVer=docker" .

FROM library/alpine:latest
COPY --from=aws-get-secret /go/src/github.com/bdwyertech/aws-get-secret/aws-get-secret /usr/local/bin/

ARG BUILD_DATE
ARG VCS_REF

LABEL org.opencontainers.image.title="bdwyertech/aws-get-secret" \
      org.opencontainers.image.version=$VCS_REF \
      org.opencontainers.image.description="For retrieving a secret from AWS Parameter Store" \
      org.opencontainers.image.authors="Brian Dwyer <bdwyertech@github.com>" \
      org.opencontainers.image.url="https://hub.docker.com/r/bdwyertech/aws-get-secret" \
      org.opencontainers.image.source="https://github.com/bdwyertech/aws-get-secret.git" \
      org.opencontainers.image.revision=$VCS_REF \
      org.opencontainers.image.created=$BUILD_DATE \
      org.label-schema.name="bdwyertech/aws-get-secret" \
      org.label-schema.description="For retrieving a secret from AWS Parameter Store" \
      org.label-schema.url="https://hub.docker.com/r/bdwyertech/aws-get-secret" \
      org.label-schema.vcs-url="https://github.com/bdwyertech/aws-get-secret.git"\
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.build-date=$BUILD_DATE

RUN apk update && apk upgrade \
    && apk add --no-cache bash ca-certificates curl \
    && adduser aws-get-secret -S -h /home/aws-get-secret

USER aws-get-secret
WORKDIR /home/aws-get-secret
CMD ["bash"]
