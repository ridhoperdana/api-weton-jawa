#!/bin/sh

echo "Starting Nginx with API_HOST=${API_HOST}"

# Replace ${API_HOST} with the actual value of the API_HOST environment variable
envsubst '${API_HOST}' < /usr/share/nginx/html/script.js > /usr/share/nginx/html/script.js.tmp && mv /usr/share/nginx/html/script.js.tmp /usr/share/nginx/html/script.js

# Start Nginx
nginx -g 'daemon off;'
