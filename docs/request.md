# List of Request made for the service

## Get Account Info

### get user account id
- used to get the id of the user from the token input to be used to get a list of models followed

[json response example](/docs/json-examples/01-get-user-id.json)
```
# GET requests to url with AUTH Token in headers
https://apiv3.fansly.com/api/v1/account/settings
```

### Get Following List
- used to get the list of creators the user follows

[json response example](/docs/json-examples/02-get-following-list.json)
```
# GET request to url with AUTH Token in headers
https://apiv3.fansly.com/api/v1/account/{user-id}/following?before=0&after=0&limit=200&offset=0
```

### Get Creator Info
- used to get the Account Name, Avatar, and Banner to be displayed on the website

[json response example](/docs/json-examples/03-get-creator-info.json)
```
# GET request to url with AUTH Token in headers
https://apiv3.fansly.com/api/v1/account?ids={model-ids}&ngsw-bypass=true
```

## Get Creator Posts

- used to get a list of post form a creator

[json response example](/docs/json-examples/04-get-creator-posts.json)
```
# GET request to url with AUTH Token in headers
https://apiv3.fansly.com/api/v1/timelinenew/{model-accountId}?before=0&after=0&wallId=&contentSearch=&ngsw-bypass=true
```