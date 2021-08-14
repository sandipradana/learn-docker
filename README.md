docker build . -t learn-docker

docker run -i -t -e INSTANCE_ID=local -e HOST=0.0.0.0 -e PORT=8080 -p 8080:8080 --name learn-docker learn-docker