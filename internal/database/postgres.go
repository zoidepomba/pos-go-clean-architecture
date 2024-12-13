package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq" // Driver para PostgreSQL
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB(user, password, host, port, dbname string) *PostgresDB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	return &PostgresDB{DB: db}
}

func (p *PostgresDB) Close() {
	if err := p.DB.Close(); err != nil {
		log.Printf("Erro ao fechar a conexão com o banco de dados: %v", err)
	}
}
