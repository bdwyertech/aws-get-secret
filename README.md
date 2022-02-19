# aws-get-secret
[![Build Status](https://github.com/bdwyertech/aws-get-secret/workflows/Go/badge.svg?branch=master)](https://github.com/bdwyertech/aws-get-secret/actions?query=workflow%3AGo+branch%3Amaster)
[![](https://images.microbadger.com/badges/image/bdwyertech/aws-get-secret.svg)](https://microbadger.com/images/bdwyertech/aws-get-secret)
[![](https://images.microbadger.com/badges/version/bdwyertech/aws-get-secret.svg)](https://microbadger.com/images/bdwyertech/aws-get-secret)

### Background
This is a simple CLI utility to retrieve a secret from AWS Parameter Store.  Sometimes you simply need to retrieve a secret.  Additionally, you might have to assume another role in order to do this, maybe because a secret resides in another account.  In either case, this lightweight utility is designed to address this need with cross-platform compatibility.


### Example Usage in GitLab CI
```yaml
release:
  stage: release
  image: golang:1.14-alpine
  variables:
    SECRET: '/sa-my-team-automation-acct'
    SECRET_ROLE_ARN: 'arn:aws:iam::123456789876:role/sa-automation-cred-retrieval'
    SECRET_ROLE_EXT_ID: $MY_ROLE_SHARED_SECRET
  before_script:
    - > # Setup aws-get-secret
      if [ ! -f .cache/aws-get-secret ]; then
        mkdir -p ${CI_PROJECT_DIR}/.cache
        wget -qO- https://github.com/bdwyertech/aws-get-secret/releases/download/v0.0.3/aws-get-secret_linux_amd64.tar.gz | tar zxf - --directory .cache
      fi
      && /bin/ln -sfn ${CI_PROJECT_DIR}/.cache/aws-get-secret /usr/local/bin/aws-get-secret
  script:
    - export MY_SECRET=$(aws-get-secret $SECRET -r $SECRET_ROLE_ARN -e $SECRET_ROLE_EXT_ID) && test -n "$MY_SECRET"
    - goreleaser release
```
