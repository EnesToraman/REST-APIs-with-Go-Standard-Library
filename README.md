# REST API design with Go Standard Library
Project is deployed to Heroku. Accessible via URL below:
[go-task-v1.herokuapp.com](https://go-task-v1.herokuapp.com/)
## Endpoints

- /users/
- /products/
- /payments/
- /brands/
- /categorires/
- /customers/
- /baskets/

### Query Parameters

1) /users/ -> id, email, username, isActive
2) /products/ -> sku, name, price, stock
3) /payments/ -> userId, amount, discount, tax
4) /brands/ -> id, name, productQty, totalWorth
5) /categories/ -> id, name, productQty, isMain
6) /customers/ -> id, userID, purchaseAmount, OrderQty
7) /baskets/ -> userId, productId, sku, quantity

#### Example - endpoints with query parameters

- /users/?id=5
- /products/?sku=11004545
- /payments/?userID=2