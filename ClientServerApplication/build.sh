# docker container up
docker rm -f prometheus
cd build/dev && docker-compose up -d
cd ../..

# build client and server
go mod tidy
rm -fr ./client/cmd/client_out
rm -fr ./server/cmd/server_out
cd ./server/cmd && go build -o server_out main.go
cd ../../
cd ./client/cmd && go build -o client_out main.go
cd ../../

# run client and server
./server/cmd/server_out & ./client/cmd/client_out