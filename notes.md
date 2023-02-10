https://pocketbase.io/

https://github.com/pocketbase/js-sdk

https://next.vuetifyjs.com/en/components/all/

https://vueuse.org/core/useEventSource/

https://github.com/bmdavis419/pocketbase-first-impressions

~~https://www.reddit.com/r/pocketbase/comments/y1fzry/serve_a_static_web_app_with_pocketbase/~~

https://github.com/pocketbase/pocketbase/blob/master/examples/base/main.go

https://rubberduckdebugging.com/



# Install Taskfile npm
npm install -g @go-task/cli

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