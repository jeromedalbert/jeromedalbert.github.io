#!/bin/bash

git add --all
git commit -v
git push

bundle exec rake generate
bundle exec rake deploy
