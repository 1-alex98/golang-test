package test

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"testing"
	"trading/core"
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image: "postgres:12.4-alpine",
		Env: map[string]string{
			"POSTGRES_DB":       "postgres",
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "banana",
		},
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForLog("ready to accept connections"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	os.Setenv("POSTGRES_USER", "postgres")
	host, _ := postgres.Host(ctx)
	os.Setenv("POSTGRES_HOST", host)
	os.Setenv("POSTGRES_PASSWORD", "banana")
	err = os.Chdir("..")
	fmt.Println(os.Getwd())
	if err != nil {
		panic(err)
	}
	core.Setup()
	code := m.Run()
	os.Exit(code)
}
