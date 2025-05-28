#!/bin/bash
Address="localhost:6985/user/"
# Helper function to print the result
print() {
  if [ $1 -eq 0 ]; then
    echo "[PASS] $2"
  else
    echo "[FAIL] $2 (HTTP $1)"
  fi
}


# 1. POST request
echo "Testing POST /api/users"
status=$(curl --location 'localhost:6985/user/create' \
         --header 'Content-Type: application/json' \
         --data-raw '{
             "firstName":"Mohsen",
             "lastName":"Taheri",
             "email":"mhthrh@gmail.com",
             "phoneNumber":"6047277989",
             "userName":"mhthrh",
             "password":"Qaz@123123"
         }')
print $status "POST /api/users"

# 2. GET request with headers
echo "Testing GET /api/users with headers"
status=$(curl --location 'localhost:6985/user/get?userName=mhthrh')
print $status "GET /api/users with headers"


# 3. PUT request
echo "Testing PUT /api/users/1"
status=$(curl --location --request PUT 'localhost:6985/user/update' \
         --header 'Content-Type: application/json' \
         --data-raw '{
             "firstName":"Mohsen",
             "lastName":"Taheri Rozbehani",
             "email":"mhthrh@gmail.com",
             "phoneNumber":"6047277989",
             "userName":"mhthrh",
             "password":"Qaz@123123"
         }')
print $status "PUT /api/users/1"


# 4. DELETE request
echo "Testing DELETE /api/users/1"
status=$(curl --location --request DELETE 'localhost:6985/user/delete?userName=mhthrh')
print $status "DELETE /api/users/1"


