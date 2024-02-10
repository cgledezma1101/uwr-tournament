#!/bin/bash
docker run --rm --volume "$PWD":/usr/src/app --workdir /usr/src/app ruby:2.7.8 bundle install
