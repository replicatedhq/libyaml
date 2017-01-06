FROM golang:1.6

ENV PROJECTPATH=/go/src/github.com/replicatedhq/libyaml

WORKDIR $PROJECTPATH

CMD ["/bin/bash"]
