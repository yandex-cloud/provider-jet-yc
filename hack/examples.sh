#!/usr/bin/env bash
set -aeuo pipefail

function join {
  local f=${1-}
  if shift; then
    printf %s "$f" "${@/#/,}"
  fi
}

## EXCLUDED ON PURPOSE:
# providerconfig gets created on initializing tests
# testing folder requires cloud admin privileges, not feasible
# alb/targetgroup needs instance IPs explicitly provided
# container/repository needs registry ID explicitly provided
# datatransfer/transfer actually runs upon creation
# message/queue works, but cleaning it in case of tests failure is a problem

## CURRENTLY FAILING:
# kubernetes/nodegroup needs investigation
# storage/object needs investigation
# storage/bucket does get removed, but Crossplane receives Forbidden; needs investigation
# securitygroup and securitygrouprule can be created, but not really altered (yet?), so don't pass tests
all=$(find ${1} -name "*.yaml" \
-not -path "*/vpc/securitygroup*.yaml" \
-not -path "*/kubernetes/nodegroup.yaml" \
-not -path "*/storage/object.yaml" \
-not -path "*/storage/bucket.yaml" \
-not -path "*/alb/*" \
-not -path "*/container/repository.yaml" \
-not -path "*/datatransfer/transfer.yaml" \
-not -path "*/message/queue.yaml" \
-not -path "*/providerconfig/providerconfig.yaml" \
-not -path "*/resourcemanager/folder.yaml" \
-not -path "*/mdb/kafkaconnector.yaml")

join $all

