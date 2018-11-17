#!/bin/bash
echo "making requests"
curl http://localhost:8080
curl -X POST http://localhost:8080 -d '{"method":"post"}'
curl -X DELETE http://localhost:8080 -d '{"method":"delete"}'
curl -X PUT http://localhost:8080 -d '{"method":"put"}'
curl -X PATCH http://localhost:8080 -d '{"method":"patch"}'

