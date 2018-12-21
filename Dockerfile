# --template:common-preamble
ARG BUILDER_IMAGE
ARG RUNTIME_IMAGE
FROM $BUILDER_IMAGE AS builder

# --template:key-provisioning-for-git-pulls-during-build
COPY docker_key /root/.ssh/id_rsa
RUN chmod 0400 /root/.ssh/id_rsa && \
  echo "Host *\n    StrictHostKeyChecking no" > /root/.ssh/config && \
    git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"

# --template:build-go-bus-microservice<REPO>
WORKDIR /go/src/envoy-sidecar
COPY go.mod Makefile . ./
RUN PROTOBUF=false make  go-unit-test go-cmds

# --template:run-time-packaging-compliance
FROM $RUNTIME_IMAGE
ARG DEPLOYMENT
LABEL "Deployment"="$DEPLOYMENT"
ARG GIT_COMMIT
LABEL "git_commit"="$GIT_COMMIT"
ARG CI_INFO
LABEL "ci_info"="$CI_INFO"

# --inject:component-level-run-time-provisioning
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates

# --template:run-time-entry-point<REPO,ENTRYPOINT>
COPY --from=builder /go/src/envoy-sidecar/bin/envoy-sidecar /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/envoy-sidecar"]
