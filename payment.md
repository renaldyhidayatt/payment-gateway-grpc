## Hello

curl -X GET http://localhost:5000/api/auth/hello


## Auth Register

curl -X POST http://localhost:5000/api/auth/register \
-H "Content-Type: application/json" \
-d '{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe@example.com",
  "password": "password123",
  "confirm_password": "password123"
}'


## Auth Login
curl -X POST http://localhost:5000/api/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "john.doe@example.com",
  "password": "password123"
}'


## Account Test

curl -X POST http://localhost:5000/api/auth/register \
-H "Content-Type: application/json" \
-d '{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john.doe2@example.com",
  "password": "password123",
  "confirm_password": "password123"
}'



## Saldo 

### Hello

curl -X GET http://localhost:5000/api/saldo/hello \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Get Saldos
curl -X GET http://localhost:5000/api/saldo/ \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Get Saldo
curl -X GET http://localhost:5000/api/saldo/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Get Saldos User

curl -X GET http://localhost:5000/api/saldo/user-all/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Get Saldo UserId

curl -X GET http://localhost:5000/api/saldo/user/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Create

curl -X POST http://localhost:5000/api/saldo/create \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "user_id": 1,
    "total_balance": 500000
}'

### Update

curl -X PUT http://localhost:5000/api/saldo/update/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "saldo_id": 1,
    "user_id": 1,
    "total_balance": 500000,
    "withdraw_amount": 5000,
    "withdraw_time": "2024-02-27T08:00:00Z"
}'

### Delete

curl -X GET http://localhost:5000/api/saldo/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


## Topup

### Hello

curl -X GET http://localhost:5000/api/topup/hello \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Find All Topup
curl -X GET http://localhost:5000/api/topup/ \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Find Topup

curl -X GET http://localhost:port/api/topup/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Find Users Toup History

curl -X GET http://localhost:5000/api/topup/user-all/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Find User Topup

curl -X GET http://localhost:5000/api/topup/user/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Create Toup

curl -X POST http://localhost:5000/api/topup/create \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NzM0MX0.BGGqLN56Nt7gE4-UpG0hphgcCq3cC2WmOeNjqewt10Y" \
-d '{
    "user_id": 1,
    "topup_no": "123456789",
    "topup_amount": 500000,
    "topup_method": "paypal"
}'


### Update Topup

curl -X PUT http://localhost:5000/api/topup/update/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NzM0MX0.BGGqLN56Nt7gE4-UpG0hphgcCq3cC2WmOeNjqewt10Y" \
-d '{
    "user_id": 1,
    "topup_id": 1,
    "topup_amount": 500000,
    "topup_method": "paypal"
}'

### Delete Topup

curl -X DELETE http://localhost:5000/api/topup/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"



### Transfer

### Hello


curl -X GET http://localhost:5000/api/transfer/hello \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Find Transfers

curl -X GET http://localhost:5000/api/transfer/ \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Find Transfer

curl -X GET http://localhost:5000/api/transfer/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"
 
### Find Transfer Users

curl -X GET http://localhost:5000/api/transfer/user-all/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Find Transfer User

curl -X GET http://localhost:5000/api/transfer/user/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


### Create Transfer


curl -X POST http://localhost:5000/api/transfer/create \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "transfer_from": 1,
    "transfer_to": 2,
    "transfer_amount": 50000
}'

### Update Transfer

curl -X PUT http://localhost:5000/api/transfer/update/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "transfer_id": 1,
    "transfer_from": 1,
    "transfer_to": 2,
    "transfer_amount": 50000
}'

### Delete Transfer

curl -X DELETE http://localhost:5000/api/transfer/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


## Withdraw 

### Hello

curl -X GET http://localhost:5000/api/withdraw/hello \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Find Withdraws

curl -X GET http://localhost:5000/api/withdraw/ \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"



### Find Withdraw
curl -X GET http://localhost:5000/api/withdraw/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"



### Find Withdraw Users

curl -X GET http://localhost:5000/api/withdraw/user-all/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"



### Find Withdraw User


curl -X GET http://localhost:5000/api/withdraw/user/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"

### Create Withdraw
curl -X POST http://localhost:5000/api/withdraw/create \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "user_id": 1,
    "withdraw_amount": 500,
    "withdraw_time": "2024-02-27T08:00:00Z"
}'

### Update Withdraw

curl -X PUT http://localhost:5000/api/withdraw/update/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI" \
-d '{
    "withdraw_id": 1,
    "user_id": 1,
    "withdraw_amount": 1000,
    "withdraw_time": "2024-02-27T08:00:00Z"
}'


### Delete Withdraw

curl -X DELETE http://localhost:5000/api/withdraw/1 \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiSm9obiBEb2UiLCJhZG1pbiI6dHJ1ZSwic3ViIjoiMSIsImV4cCI6MTcwOTI2NDU4MX0.B7kubMnPJ8f4xNGtsp_LnidfPcJiJUhrc7uXx5FZ6mI"


