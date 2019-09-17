#!/bin/bash
#tag=v.$(date +'%Y%m%d.%H%M').$(git rev-parse --short HEAD)
tag=v.$(date +'%Y.%m%d.%H%M')
git push
git tag "$tag"
git push origin "$tag"
