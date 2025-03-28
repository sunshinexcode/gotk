docker rmi $(docker images $1 | awk '{print $3}' | sed '1,3d')
docker images | grep $1