#!/bin/sh

# Add the liveness command to start
liveness --resource http://localhost:80 &

# Run the included nginx entrypoint
/docker-entrypoint.sh $@

# Run nginx
nginx -g 'daemon off;'