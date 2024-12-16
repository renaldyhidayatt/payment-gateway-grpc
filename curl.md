## Auth

### Login
```sh
curl -X POST http://172.24.0.3:5000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "password": "securepassword"
  }'
```


### Register

```sh
curl -X POST http://172.24.0.3:5000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword",
    "confirm_password": "securepassword"
  }'
```

--------------------------------------------------

## User

### FindAll
```sh
curl -X GET "http://localhost:5000/findall" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>"
```

### FindById
```sh
curl -X GET http://localhost:5000/find/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>"
```


### FindByActive
```sh
curl -X GET http://localhost:5000/active \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>"
```

### FindByTrashed

```sh
curl -X GET http://localhost:5000/trashed \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>"
```

### Create

```sh
curl -X POST http://localhost:5000/user/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.doe@example.com",
    "password": "securepassword",
    "confirm_password": "securepassword"
}'
```

### Update
```sh
curl -X POST http://localhost:5000/user/update/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{
    "user_id": 1,
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.updated@example.com",
    "password": "newsecurepassword",
    "confirm_password": "newsecurepassword"
}'
```


### Trashed

```sh
curl -X POST http://localhost:5000/user/trashed/1 \
  -H "Authorization: Bearer <your_token>"
```

### Restore

```sh
curl -X POST http://localhost:5000/user/restore/1 \
  -H "Authorization: Bearer <your_token>"
```

### Permanent

```sh
curl -X DELETE http://localhost:5000/user/permanent/1 \
  -H "Authorization: Bearer <your_token>"
```

--------------------------------------------------