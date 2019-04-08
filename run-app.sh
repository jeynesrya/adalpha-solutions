# backend services
docker-compose up -d elasticsearch kibana postgres
sleep 5
docker build -t adalpha .
docker-compose up -d backend
# user interface
cd ui
npm install
npm run serve
