FROM golang:1.5

ENV PROJECTPATH=/go/src/github.com/replicatedhq/libyaml

WORKDIR $PROJECTPATH

CMD ["/bin/bash"]
