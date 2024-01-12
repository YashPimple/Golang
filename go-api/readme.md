# GO-API using gin
GET Request: curl localhost:8080/books
POST Request: curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
PATCH request: curl localhost:8080/checkout?id=2 --request "PATCH"