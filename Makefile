run-db:
	docker run -d --name demo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-root-pass -e MYSQL_DATABASE=todo_db mysql:8.0

stop-db:
	docker stop demo-mysql
	docker rm demo-mysql

build-run:
	go build -o bin/todo cmd/main.go
	./bin/todo

test:
	go test -v ./...

run:
	go run cmd/main.go

connect-db:
	docker exec -it demo-mysql mysql -uroot -pmy-root-pass todo_db