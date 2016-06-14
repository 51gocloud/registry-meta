docker stop meta
docker rm $(docker ps -qa)
docker rmi registry-meta:latest
docker build --no-cache --rm=true -t registry-meta:latest .
docker run --name meta -d -p 6000:6000 registry-meta:latest registry-meta
