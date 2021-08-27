#!/bin/bash

set -e

GV="devops:v1alpha1"

rm -rf ./pkg/client
./hack/generate_group.sh "client,lister,informer" github.com/hchenc/pager/pkg/client github.com/hchenc/pager/pkg/apis "$GV" --output-base=./  -h "$PWD/hack/boilerplate.go.txt"
mv github.com/hchenc/pager/pkg/client ./pkg/
rm -rf ./github.com/hchenc