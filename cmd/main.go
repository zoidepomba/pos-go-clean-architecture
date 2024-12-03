package main

import (
"database/sql"
"fmt"
"log"
"net"
"net/http"
"google.golang.org/grpc"

_ "github.com/lib/pq" // Driver para PostgreSQL

httpHandler "project/internal/delivery/http"
grpcHandler "project/internal/delivery/grpc"
"project/internal/repository"
"project/internal/service"

pb "project/proto"

)

func main() {
// Configuração do banco de dados
db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/orders_db?sslmode=disable")
if err != nil {
	log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
}
defer db.Close()

// Inicializar repositório
orderRepo := &repository.OrderRepository{DB: db}

// Inicializar serviço
orderService := &service.OrderService{Repo: orderRepo}

// Inicializar HTTP Handler
orderHandler := &httpHandler.OrderHandler{Service: orderService}

// Inicializar servidor HTTP
go func() {
	http.HandleFunc("/order", orderHandler.ListOrders)
	fmt.Println("Servidor HTTP rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
}()

// Inicializar servidor GRPC
go func() {
	grpcServer := grpc.NewServer()
	orderGrpcService := &grpcHandler.OrderService{Service: orderService}
	pb.RegisterOrderServiceServer(grpcServer, orderGrpcService)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Erro ao iniciar servidor GRPC: %v", err)
	}
	fmt.Println("Servidor GRPC rodando na porta 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Erro ao rodar servidor GRPC: %v", err)
	}
}()

// Inicializar servidor GraphQL (opcional)
// Aqui você pode usar bibliotecas como gqlgen ou graphql-go para configurar o servidor GraphQL.

// Manter a aplicação rodando
select {}
}