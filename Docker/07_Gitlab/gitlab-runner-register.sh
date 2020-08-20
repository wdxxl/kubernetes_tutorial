#!/bin/sh
# Get the registration token from:
# http://localhost:10080/mygroup/myproject/-/settings/ci_cd

registration_token=cQP4MsbejC8jonhP4221

docker exec -it gitlab-runner1 \
  gitlab-runner register \
    --non-interactive \
    --registration-token ${registration_token} \
    --locked=false \
    --description shell-runner \
    --url http://gitlab \
    --executor shell