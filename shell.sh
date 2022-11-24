# Redis with docker
docker container create --name smallapp -p 6379:6379 --cpus 2 --memory=50m redis
docker container start smallapp
