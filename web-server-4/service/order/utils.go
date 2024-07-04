package order

import (
	"fmt"

	"github.com/jmavila/golang/web-server-4/models"
)

func GetCheckoutProductsIDs(items []models.CheckoutItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductId)
		}
		productIds[i] = item.ProductId
	}

	return productIds, nil
}

func (h *Handler) createOrders(ps []models.Product, items []models.CheckoutItem, userID int) (int, float64, error) {
	productMap := make(map[int]models.Product)
	for _, product := range ps {
		productMap[product.ID] = product
	}
	if err := checkIfProductIsInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	totalPrice := calculateTotalPrice(items, productMap)

	for _, item := range items {
		product := productMap[item.ProductId]
		product.Quantity -= item.Quantity
		h.productStore.UpdateProduct(&product)
	}

	orderId, err := h.store.CreateOrder(models.Order{
		UserId:  userID,
		Total:   totalPrice,
		Status:  "pending",
		Address: "todo",
	})
	if err != nil {
		return 0, 0, err
	}

	return orderId, totalPrice, nil
}
func calculateTotalPrice(items []models.CheckoutItem, products map[int]models.Product) float64 {
	var total float64
	for _, item := range items {
		product := products[item.ProductId]
		total += product.Price
	}
	return total
}

func checkIfProductIsInStock(items []models.CheckoutItem, products map[int]models.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("car is empty")
	}
	for _, item := range items {
		product, ok := products[item.ProductId]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductId)
		}
		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %d is not available in the quantity requested", product.Name)
		}
	}
	return nil
}
