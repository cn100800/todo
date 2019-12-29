#!/bin/bash
git tag |xargs git push origin --delete #删除远程tag
git tag |xargs git tag -d #删除本地tag
