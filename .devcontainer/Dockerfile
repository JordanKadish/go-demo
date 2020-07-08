FROM golang:latest
COPY . /workspace
RUN apt-get update \
    && export GO111MODULE=on \
    && apt-get -y install --no-install-recommends apt-utils dialog 2>&1 \
    && go get -v golang.org/x/tools/gopls@latest \
    # Docker CLI
    && apt-get -y install apt-transport-https ca-certificates curl gnupg-agent software-properties-common \
    && curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | (OUT=$(apt-key add - 2>&1) || echo $OUT) \
    && add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" \
    && apt-get update \
    && apt-get install -y docker-ce-cli
