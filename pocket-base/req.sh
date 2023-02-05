# Run PocketBase Docker
docker run -d -p 8090:8090 --name pocketbase pocketbase/pocketbase

# Run PocketBase Default Port
go run . serve

# Set Initial Admin
curl --request POST \
    --url http://127.0.0.1:8090/api/admins \
    --header 'Content-Type: application/json' \
    --data '{"email": "test@email.com", "password": "P@ssw0rd***", "passwordConfirm": "P@ssw0rd***"}'

# Login PocketBase Admin
curl --request POST \
  --url http://127.0.0.1:8090/api/admins/auth-with-password \
  --header 'Content-Type: application/json' \
  --data '{"identity": "test@email.com","password": "P@ssw0rd***"}'