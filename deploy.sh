#!/bin/sh -eu

REPOSITORY=docker.io/fillmorelabs
PROFILE=fillmore-labs

IMAGE=$(\
  env KO_DOCKER_REPO=$REPOSITORY/login-sample\
  ko build --bare --sbom none\
)

cat << _IMAGE > k8s/$PROFILE/image.yaml
---
apiVersion: builtin
kind: ImageTagTransformer
metadata:
  name: login-sample
imageTag:
  name: login-sample-image
  newName: ${IMAGE%%@*}
  digest: ${IMAGE#*@}
_IMAGE

sops -d k8s/$PROFILE//secrets.enc.env > k8s/$PROFILE/secrets.env

kubectl apply -k k8s/$PROFILE

rm -f k8s/$PROFILE/image.yaml k8s/$PROFILE/secrets.env
