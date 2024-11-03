run-db:
	docker run -d --name demo-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-root-pass -e MYSQL_DATABASE=todo_db mysql:8.0