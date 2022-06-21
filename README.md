Como crear/aplicar migraciones
https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md

$ migrate -source file://path/to/migrations -database postgres://localhost:5432/

### ejemplo creacion

migrate create -ext sql -dir migrations -seq create_users_table

### ejemplo migracion
migrate -database postgres://postgres:1234@localhost:5432/go_prueba?sslmode=disable -path migrations up

### comando para desarrollo
CompileDaemon --build="swag init && go build ./main.go" --directory=./ --exclude-dir=./docs --command="./main"
