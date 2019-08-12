# !/bin/bash

export IMAGE_TAG=$(cat VERSION)
export DOCKER_CLI_EXPERIMENTAL=enabled
export IMAGE_SUFFIXES=(base voltha ofagent netconf shovel onos tester config-push dashd cli portainer envoy alarm-generator test_runner)

for SUFFIX in ${IMAGE_SUFFIXES[@]}; do
  IMAGE_NAME="voltha-${SUFFIX}"
  docker manifest create --amend cachengo/$IMAGE_NAME:$IMAGE_TAG cachengo/$IMAGE_NAME-x86_64:$IMAGE_TAG cachengo/$IMAGE_NAME-aarch64:$IMAGE_TAG
  docker manifest push --purge cachengo/$IMAGE_NAME:$IMAGE_TAG
done
