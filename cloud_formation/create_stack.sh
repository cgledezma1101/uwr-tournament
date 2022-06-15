#!/bin/bash
if [[ $# -ne 4 ]]; then
    echo "Illegal number of arguments: $#. Usage is:"
    echo "$0 database-snapshot-arn database-password sendgrid-password alb-certificate-arn"
    exit 1
fi

aws cloudformation create-stack \
    --region us-east-1 \
    --stack-name uwr-tournaments \
    --template-body file://webserver_cloudformation.yml \
    --parameters "[
        { \"ParameterKey\": \"TournamentsSnapshotArn\", \"ParameterValue\": \"$1\" },
        { \"ParameterKey\": \"DatabasePassword\", \"ParameterValue\": \"$2\" },
        { \"ParameterKey\": \"SendgridPassword\", \"ParameterValue\": \"$3\" },
        { \"ParameterKey\": \"AlbCertificateArn\", \"ParameterValue\": \"$4\" }
    ]"