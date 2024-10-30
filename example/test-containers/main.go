package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func main() {
	ctx := context.Background()

	// 启动 PostgreSQL 容器
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_USER":     "user",
			"POSTGRES_DB":       "graphnode",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer postgresC.Terminate(ctx)

	// 获取容器 IP 和端口
	host, err := postgresC.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal(err)
	}

	// 连接到 PostgreSQL 数据库
	dsn := fmt.Sprintf("host=%s port=%s user=user password=password dbname=graphnode sslmode=disable", host, port.Port())
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 获取数据库名称
	var dbName string
	err = db.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to database: %s\n", dbName)
}
