#!/bin/bash
if [[ $# -ne 1 ]]; then
    echo "Illegal number of arguments: $#. Usage is:"
    echo "$0 database-snapshot-arn"
    exit 1
fi

aws cloudformation create-stack \
    --region us-east-1 \
    --stack-name uwr-tournaments \
    --template-body file://webserver_cloudformation.yml \
    --capabilities CAPABILITY_NAMED_IAM \
    --parameters "[
        { \"ParameterKey\": \"TournamentsSnapshotArn\", \"ParameterValue\": \"$1\" }
    ]"

