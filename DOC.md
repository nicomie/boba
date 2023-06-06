# API

## USERS
### User
- Name: String
- ID: Integer
- Favorites: Order[]
- Last_order: Order

### Endpoints
- GET: /users
- GET: /users/id
- POST: /users
- POST /users/id/favorites
- POST /users/id/last_order

## ORDERS
### Order
- ID: User.id
- Size: Enum
- Flavors: Enum[]
- Sweetness: Enum
- Milk: Enum

### Endpoints
- GET: /orders
- GET: /orders/id
- POST: /orders


## 