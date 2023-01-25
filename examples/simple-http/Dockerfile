FROM nginx:latest

RUN apt update && apt install curl

ARG LIVENESS_VERSION=0.1.0
# Download the liveness release tarball for our LIVENESS_VERSIOn and extract the
# binary into /usr/bin. The results of uname -m must be passed through sed since
# it returns aarch64 but the release tarballs use arm64.
RUN curl -o /tmp/liveness.tar.gz -L https://github.com/chasinglogic/liveness/releases/download/${LIVENESS_VERSION}/liveness_Linux_$(uname -m | sed s/aarch64/arm64/).tar.gz && \
    tar -C /usr/bin -xzvf /tmp/liveness.tar.gz liveness

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]