FROM golang:1.11.4

LABEL com.github.actions.name="Go"
LABEL com.github.actions.description="Pass go processes"
LABEL com.github.actions.icon="check-square"
LABEL repository="https://github.com/QuentinBrosse/pass"
LABEL homepage="https://github.com/QuentinBrosse/pass"
LABEL maintainer="Quentin Brosse <quentin.brosse@icloud.com>"

RUN go get golang.org/x/tools/cmd/goimports

ENV GO111MODULE on

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
