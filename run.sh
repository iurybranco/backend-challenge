#IT INSTALLS CURL IF DOES NOT EXIST IN THE MACHINE
if curl -V; then echo "curl installed"; else sudo apt-get update && sudo apt-get -y install curl; fi
#IT GETS UP ALL THE CONTAINERS
docker-compose up -d
#IT CHECKS IF MONGODB IS UP
until curl --connect-timeout 10 --silent --show-error localhost:27017 | grep 'It looks like you are trying to access MongoDB over HTTP on the native driver port.'; do echo "waiting"; sleep 2; done
#IT RUNS db_seed.sh INSIDE OF MONGO CONTAINER
docker-compose exec mongodb sh db_seed.sh
