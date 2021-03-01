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
    "data": {
        "items": [
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

### Get Item

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
