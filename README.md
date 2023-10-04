# caddy-web-proxy
A simple proxy server implemented with Caddy

# Build

```shell
docker buildx build \
-t saladtechnologies/caddy-web-proxy:latest \
--provenance=false \
--output type=docker \
.
```

# Run

```shell
docker run \
-p 3000:3000 \
saladtechnologies/caddy-web-proxy:latest
```

# Use

To access https://salad.com/salad-cloud/generative-ai, use the following command:

```shell
curl  -X GET \
  'https://localhost:3000/salad-cloud/generative-ai' \
  --header 'Accept: */*' \
  --header 'x-upstream-host: salad.com'
```

Note, all parts of the request are proxied as-is, including the headers. The only exception is the `Host` header, which is replaced with the value of the `x-upstream-host` header.