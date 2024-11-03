run-db:
	docker run -d --name demo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-root-pass -e MYSQL_DATABASE=todo_db mysql:8.0

stop-db:
	docker stop demo-mysql
	docker rm demo-mysql

build:
	go build -o bin/todo cmd/main.go

test:
	go test -v ./...

run: build
	./bin/todo