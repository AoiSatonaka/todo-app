FROM sapmachine:lts-ubuntu-jammy

ARG NODE_MAJOR=20

RUN apt update && \
  apt install -y ca-certificates curl gnupg && \
  mkdir -p /etc/apt/keyrings && \
  # node source gpgkey import
  curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg && \
  # create deb repository
  echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list && \
  # install Nodejs
  apt update && apt install -y nodejs && \
  # install firebase cli
  npm i -g firebase-tools

# start firebase local emulators
ENV GCLOUD_PROJECT=yourProjectId
CMD firebase emulators:start --import ./firebase/data --export-on-exit
