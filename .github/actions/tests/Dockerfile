FROM golang:1.11.4

LABEL com.github.actions.name="Tests"
LABEL com.github.actions.description="Runs go tests"
LABEL com.github.actions.icon="check-square"
LABEL repository="https://github.com/QuentinBrosse/pass"
LABEL homepage="https://github.com/QuentinBrosse/pass"
LABEL maintainer="Quentin Brosse <quentin.brosse@icloud.com>"

ENV GO111MODULE on

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
