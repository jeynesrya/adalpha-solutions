docker-compose -f docker-compose-test.yml up -d
sleep 5
go test ./... -v
docker-compose -f docker-compose-test.yml down
