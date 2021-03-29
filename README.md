# GERVICHSTORE - Backend

Backend server for gervichstore-client website.

## API SPEC

---

## Users

### Login User

-   Method : `GET`
-   Endpoint : `/api/users/login`
-   Header : 
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :
```json
{
    "email": "String",
    "password": "String" 
}
```
-   Response :
```json
{
    "code": "Number",
    "token": "String"
}
```

### Register User

-   Method : `POST`
-   Endpoint : `/api/users/register`
-   Header : 
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :
```json
{
    "phoneNumber": "String",
    "password": "String",
    "firstName": "String",
    "lastName": "String",
}
```
-   Response : 
```json
{
    "code": "Number"
}
```

### Update User

-   Method : `PUT`
-   Endpoint : `/api/users`
-   Header : 
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Body :
```json
{   
    "phoneNumber": "String",
    "password": "String",
    "firstName": "String",
    "lastName": "String",
    "email": "String",
    "gender": "String",
    "image": "String"
}
```
-   Response : 
```json
{
    "code": "Number"
}
```

---


## Products

### Get All Products

-   Method : `GET`
-   Endpoint : `/api/products`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number",
    "data": [
        {
            "id": "Number",
            "image": "String",
            "name": "String",
            "price": "Number",
            "stock": "Number",
            "category": "String",
            "created_at": "TimeStamp",
            "updated_at": "TimeStamp"
        }
    ]
}
```

### Create Product

-   Method : `POST`
-   Endpoint : `/api/products/`
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

### Get Product By Id

-   Method : `GET`
-   Endpoint : `/api/products/:id_product`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
-   Response :

```json
{
    "code": "Number",
    "data": {
        "id": "Number",
        "image": "String",
        "name": "String",
        "price": "Number",
        "stock": "Number",
        "category": "String",
        "created_at": "TimeStamp",
        "updated_at": "TimeStamp"
    }
}
```

### Search Product

- Method : `GET`
- Endpoint : `/api/products/:product_name`
-   Header :
    -   Content-Type : `application/json`
    -   Accept : `application/json`
- Response :
```json
{
    "code": "Number",
    "data": {
        "id": "Number",
        "image": "String",
        "name": "String",
        "price": "Number",
        "stock": "Number",
        "category": "String",
        "created_at": "TimeStamp",
        "updated_at": "TimeStamp"
    }
}
```

### Update Product

-   Method : `PUT`
-   Endpoint : `/api/products/:id_product`
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
    "category": "String",
}
```

-   Response :

```json
{
    "code": "Number"
}
```

### Delete Product

-   Method : `PUT`
-   Endpoint : `/api/products/:id_product`
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
            "products": "[]Products",
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
    "products": "[]Products",
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