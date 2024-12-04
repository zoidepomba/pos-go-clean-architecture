package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"

	_ "github.com/lib/pq" // Driver para PostgreSQL

	grpcHandler "project/internal/delivery/grpc"
	httpHandler "project/internal/delivery/http"
	"project/internal/repository"
	"project/internal/service"

	pb "project/proto"
)

func main() {
	// Pegue as variáveis de ambiente
	
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construa a string de conexão
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	log.Println(dsn)
	// Configuração do banco de dados

	db, err := sql.Open("postgres", dsn)
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
		mux := http.NewServeMux()

		// Rota para listar orders (GET)
		mux.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				orderHandler.ListOrders(w, r)
			} else if r.Method == http.MethodPost {
				orderHandler.CreateOrder(w, r)
			} else {
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			}
		})

		fmt.Println("Servidor HTTP rodando na porta 8080...")
		if err := http.ListenAndServe(":8080", mux); err != nil {
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
