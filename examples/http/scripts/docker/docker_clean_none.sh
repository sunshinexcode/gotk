docker rmi $(docker images | grep "none" | awk '{print $3}')
docker images | grep none