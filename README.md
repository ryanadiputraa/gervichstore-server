# GERVICHSTORE - Backend

Backend server for gervichstore-client website.

## API SPEC

---

## Items

### Get All Items

-   Method : `GET`
-   Endpoint : `/api/items`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number",
    "data": [
        {
            "id": "String",
            "image": "String",
            "name": "String",
            "price": "Number",
            "stock": "Number",
            "category": "String"
        }
    ]
}
```

### Create Item

-   Method : `POST`
-   Endpoint : `/api/items/`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :

```json
{
    "image": "String",
    "name": "String",
    "price": "Number",
    "stock": "Number",
    "category": "String"
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Get Item By Id

-   Method : `GET`
-   Endpoint : `/api/items/:id_item`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number",
    "data": {
        "id": "String",
        "image": "String",
        "name": "String",
        "price": "Number",
        "stock": "Number",
        "category": "String"
    }
}
```

### Search Item

- Method : `GET`
- Endpoint : `/api/items/:item_name`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
- Response :
```json
{
    "code": "Number",
    "data": {
        "id": "String",
        "image": "String",
        "name": "String",
        "price": "Number",
        "stock": "Number",
        "category": "String"
    }
}
```

### Update Item

-   Method : `PUT`
-   Endpoint : `/api/items/:id_item`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :

```json
{
    "image": "String",
    "name": "String",
    "price": "Number",
    "stock": "Number",
    "category": "String"
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Delete Item

-   Method : `PUT`
-   Endpoint : `/api/items/:id_item`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number"
}
```

---

## Transaction

### Get Transaction

-   Method : `GET`
-   Endpoint : `/api/transaction`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   x-access-token: `token`
-   Response :

```json
{
    "code": "Number",
    "data": [
        {
            "id": "String",
            "status": "String",
            "items": "[]Items",
            "totalPrice": "Number"
        }
    ]
}
```

### Create Transaction

-   Method : `POST`
-   Endpoint : `/api/transaction`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   x-access-token: `token`
-   Body :
```json
{
    "status": "String",
    "items": "[]Items",
    "totalPrice": "Number"
}
```
- Response : 
```json
{
    "code": "Number"
}
```

### Done Transaction

-   Method : `PUT`
-   Endpoint : `/api/transaction/done/:id_transaction`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   x-access-token: `token`
- Response : 
```json
{
    "code": "Number"
}
```

### Cancel Transaction

-   Method : `PUT`
-   Endpoint : `/api/transaction/cancel/:id_transaction`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
    -   x-access-token: `token`
- Response : 
```json
{
    "code": "Number"
}
```