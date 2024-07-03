package main

import (
	"context"
	"fmt"

	"github.com/albertosfernandes/api-go-rest/database"
	"github.com/albertosfernandes/api-go-rest/routes"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {

	// models.Vms = []models.VM{
	// 	{Id: 1, Name: "vm-teste01", Vcpu: 8, Mem: 12000,
	// 		Network: models.Network{Hostname: "net-01", IPAddress: "192.168.0.101", Gateway: "192.168.0.1"},
	// 		Storage: models.Storage{Name: "storage-01", Size: 100000, Type: "ssd"}},
	// 	{Id: 2, Name: "vm-teste02", Vcpu: 8, Mem: 12000,
	// 		Network: models.Network{Hostname: "net-02", IPAddress: "192.168.0.102", Gateway: "192.168.0.1"},
	// 		Storage: models.Storage{Name: "storage-02", Size: 100000, Type: "ssd"}},
	// }

	// Iniciando configuração do otel
	// 1. Criando o contexto
	ctx := context.Background()

	// 2. Configura o exportador apontando para o endpoint do Collector
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))
	if err != nil {
		fmt.Printf("Falha para configurar o exportador: %v\n", err)
		return
	}

	// 3. Cria o provedor de traces
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.Default()),
	)
	defer func() { _ = tp.Shutdown(ctx) }()

	// 4. Define o provedor de traces global
	otel.SetTracerProvider(tp)

	// 5. Obtem um Tracer
	tracer := otel.Tracer("api-go-rest-example")

	// 6. Cria um span
	ctx, span := tracer.Start(ctx, "database-connection")
	fmt.Println("Inicio Spam database")
	database.Dbconnect()

	println("Servidor inicializado, ouvindo em :8000")

	routes.HandleRequest()
	span.End()
	fmt.Println("Fim Spam database")
}
