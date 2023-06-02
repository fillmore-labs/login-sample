#!/bin/sh -eu

PROFILE=fillmore-labs

export KO_DATA_PATH=kodata/
export OAUTH_DOMAIN=fillmore-labs.eu.auth0.com
export OAUTH_CALLBACK_URL=http://127.0.0.1:8080/callback

sops exec-env k8s/$PROFILE/secrets.enc.env "go run main.go"
