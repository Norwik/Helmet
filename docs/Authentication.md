### Authentication


### Basic Authentication


### API Key Based


### OAuth 2.0

To create oauth access token using client credentials (client id and client secret)

```zsh
export APIGW_URL=http://127.0.0.1:8000
export CLIENT_ID=d221bfbc-0e47-4f92-86d2-d01fa400e7fc
export CLIENT_SECRET=d221bfbc-0e47-4f92-86d2-d01fa400e7fr

$ curl -X POST --user $CLIENT_ID:$CLIENT_SECRET $APIGW_URL/apigw/token -d 'grant_type=client_credentials'
```
