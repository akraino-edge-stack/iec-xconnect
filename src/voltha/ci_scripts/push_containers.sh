# !/bin/bash

export IMAGE_TAG=$(cat VERSION)
export AARCH=`uname -m`
export IMAGE_SUFFIXES=(base voltha ofagent netconf shovel onos tester config-push dashd cli portainer envoy alarm-generator test_runner)
export VOLTHA_BUILD=docker

make seba

for SUFFIX in ${IMAGE_SUFFIXES[@]}; do
  IMAGE_NAME="voltha-${SUFFIX}"
  docker tag $IMAGE_NAME:latest cachengo/$IMAGE_NAME-$AARCH:$IMAGE_TAG
  docker push cachengo/$IMAGE_NAME-$AARCH:$IMAGE_TAG
done
