# aws-get-secret
[![Build Status](https://github.com/bdwyertech/aws-get-secret/workflows/Go/badge.svg?branch=master)](https://github.com/bdwyertech/aws-get-secret/actions?query=workflow%3AGo+branch%3Amaster)
[![](https://images.microbadger.com/badges/image/bdwyertech/aws-get-secret.svg)](https://microbadger.com/images/bdwyertech/aws-get-secret)
[![](https://images.microbadger.com/badges/version/bdwyertech/aws-get-secret.svg)](https://microbadger.com/images/bdwyertech/aws-get-secret)

### Background
This is a simple CLI utility to retrieve a secret from AWS Parameter Store.  Sometimes you simply need to retrieve a secret.  Additionally, you might have to assume another role in order to do this, maybe because a secret resides in another account.  In either case, this lightweight utility is designed to address this need with cross-platform compatibility.
