package database

import (
	"context"

	"github.com/cyan-store/hook/log"
)

func CreateOrder(
	id, productID, transactionID,
	userID, status, quantity string,
	amount int, email, shipping,
	shippingName, city, country,
	line1, line2, postal,
	state string,
) error {
	c, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	_, err := Conn.ExecContext(c,
		`
		INSERT INTO Orders (
			id, productID, transactionID,
			userID, status, quantity,
			amount, email, shipping,
			shippingName, city, country,
			line1, line2, postal, state,
			createdAt, updatedAt
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,
			CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		)
	`,
		id, productID, transactionID,
		userID, status, quantity,
		amount, email, shipping,
		shippingName, city, country,
		line1, line2, postal, state,
	)

	defer cancel()

	if err != nil {
		log.Error.Println("[CreateOrder] Could not create order:", err)
		return err
	}

	return nil
}
