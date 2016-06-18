docker stop meta
docker stop redis
docker rm $(docker ps -qa)
docker rmi registry-meta:latest
docker build --no-cache --rm=true -t registry-meta:latest .
docker run -d -p 6379:6379 --name redis redis:latest
docker run --name meta -d --net=host -P registry-meta:latest registry-meta
