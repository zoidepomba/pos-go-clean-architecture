package database

import (
	"fmt"
	"log"
)

func RunMigration(db *PostgresDB) error {
	query := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_name VARCHAR(255) NOT NULL,
		total_amount NUMERIC(10, 2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	log.Println("Executando migration...")
	log.Println(&db)
	_, err := db.DB.Exec(query)
	if err != nil {
		log.Printf("Query de migration falhou: %s", query)
		return fmt.Errorf("erro ao executar a migration: %w", err)
	}

	log.Println("Migration executada com sucesso!")
	return nil
}
