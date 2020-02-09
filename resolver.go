package go_orders_graphql_api

import (
    "context"

    "github.com/jinzhu/gorm"
    "github.com/soberkoder/go-orders-graphql-api/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
    DB *gorm.DB
}

func (r *Resolver) Mutation() MutationResolver {
    return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func mapItemsFromInput(itemsInput []*ItemInput) []models.Item {
    var items []models.Item
    for _, itemInput := range itemsInput {
        items = append(items, models.Item{
            ProductCode: itemInput.ProductCode,
            ProductName: itemInput.ProductName,
            Quantity:    itemInput.Quantity,
        })
    }
    return items
}

func (r *mutationResolver) CreateOrder(ctx context.Context, input OrderInput) (*models.Order, error) {
    order := models.Order{
        CustomerName: input.CustomerName,
        OrderAmount:  input.OrderAmount,
        Items:        mapItemsFromInput(input.Items),
    }
    err := r.DB.Create(&order).Error
    if err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *mutationResolver) UpdateOrder(ctx context.Context, orderID int, input OrderInput) (*models.Order, error) {
    updatedOrder := models.Order{
        ID:           orderID,
        CustomerName: input.CustomerName,
        OrderAmount:  input.OrderAmount,
        Items:        mapItemsFromInput(input.Items),
    }
    err := r.DB.Save(&updatedOrder).Error
    if err != nil {
        return nil, err
    }
    return &updatedOrder, nil
}

func (r *mutationResolver) DeleteOrder(ctx context.Context, orderID int) (bool, error) {
    r.DB.Where("order_id = ?", orderID).Delete(&models.Item{})
    r.DB.Where("id = ?", orderID).Delete(&models.Order{})
    return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Orders(ctx context.Context) ([]*models.Order, error) {
    var orders []*models.Order
    err := r.DB.Preload("Items").Find(&orders).Error
    if err != nil {
        return nil, err
    }

    return orders, nil
}

/*
curl -H 'Content-Type: application/json' \
-d '{"query": "{orders{id  items {productCode quantity}}}"}' \
-X POST http://localhost:63932/query
*/
