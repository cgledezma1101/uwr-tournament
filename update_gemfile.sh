#!/bin/bash
docker run --rm --volume "$PWD":/usr/src/app --workdir /usr/src/app ruby:3.3.0 bundle install
