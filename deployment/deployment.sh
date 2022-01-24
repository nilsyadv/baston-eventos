
# login to docker repository with username & password.
cat ~/learn-tech/docker-pwd.txt | docker login --username=nileshyadv --password-stdin
uuid=$(dbus-uuidgen)

# building docker image from docker file & tag
docker build . -f deployment/dockerfile --build-arg ssh_prv_key="$(cat ~/.ssh/id_ed25519)" --tag nileshyadv/baston-eventos:$uuid

# tagged latest build image to latest
docker tag nileshyadv/baston-eventos:$uuid nileshyadv/baston-eventos:latest

# push to docker hub
docker push nileshyadv/baston-eventos:latest