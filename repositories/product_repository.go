package repositories

import (
	"context"
	"log"
	"store-api/config"
	"store-api/models"
)

func GetAllProducts() ([]models.Product, error) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, name, price FROM products")
	if err != nil {
		log.Printf("Error querying products: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			log.Printf("Error scanning product: %v\n", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
