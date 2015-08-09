#!/bin/sh
read tags
git tag -a $tags -m $tags
git push origin --tag $tags
sleep 50
git push origin --tag :$tags
git tag -d $tags