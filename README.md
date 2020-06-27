# go-orders-graphql-api
A simple GraphQL API with MySQL and GORM

To run the app, navigate to your project directory, and run the following command:
```go
go run server/server.go
```

Navigate to [http://localhost:8080/](http://localhost:8080/) on your browser to open GraphQL playground

**Create Order**
```graphql
mutation createOrder ($input: OrderInput!) {
  createOrder(input: $input) {
    id
    customerName
    items {
      id
      productCode
      productName
      quantity
    }
  }
}
```
Query Variables:
```json
{
  "input": {
    "customerName": "Leo",
    "orderAmount": 9.99,
    "items": [
      {
      "productCode": "2323",
      "productName": "IPhone X",
      "quantity": 1
      }
    ]
  }
}
```
***
**Get Orders**
```graphql
query orders {
  orders {
    id  
    customerName
    items {
      productName
      quantity
    }
  }
}
```
***
**Update Order**
```graphql
mutation updateOrder ($orderId: Int!, $input: OrderInput!) {
  updateOrder(orderId: $orderId, input: $input) {
    id
    customerName
    items {
      id
      productCode
      productName
      quantity
    }
  }
}
```
Query variables:
```json
{
  "orderId":1,
  "input": {
    "customerName": "Cristiano",
    "orderAmount": 9.99,
    "items": [
      {
      "productCode": "2323",
      "productName": "IPhone X",
      "quantity": 1
      }
    ]
  }
}
```
***
**Delete Order**
```graphql
mutation deleteOrder ($orderId: Int!) {
  deleteOrder(orderId: $orderId)
}
```
Query variables:
```json
{
  "orderId": 3
}
```

You can also use tools like [Insomnia](https://insomnia.rest/graphql/) and 
[Altair](https://altair.sirmuel.design/) to try these requests out.

***
### Tutorial

You can find the tutorial for this application at the [SoberKoder](https://www.soberkoder.com/) blog.

https://www.soberkoder.com/go-graphql-api-mysql-gorm/
