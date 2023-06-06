# API

| HTTP Status Code | Description |
| ---------------- | ----------- |
| 200 OK           | Successful request |
| 201 Created      | Created |
| 400 Bad request  | Bad input parameters |
| 403 Forbidden    | Invalid Auth token |
| 404 Not Found    | Resource location not found |

## USERS
### User
```
User {
     Name: String
     ID: Integer
     Favorites: Order[]
     Last_order: Order
 }
```
### Endpoints
- GET: /users
- GET: /users/id
- POST: /users
- POST: /users/id/favorites
- POST: /users/id/last_order

## ORDERS
### Order
```
Order {
    ID: User.id
    Size: Enum
    Flavors: Enum[]
    Sweetness: Enum
    Milk: Enum
```

### Endpoints
- GET: /orders
- GET: /orders/id
- POST: /orders


## 