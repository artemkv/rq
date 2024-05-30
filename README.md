## rq

**Request**: make your own command line for any REST API!

**R.I.P. Postman!**

- Make requests
- Replace placeholders with parameters
- Refer to env files with defaults
- Override from command line
- Chain requests like a boss!
- Feed request outputs to other request inputs


Example:

```
rq getobject objectid=12345 -e dev
```

This looks for the `getobject` scenario in a scenario file (`rq.json`):

```
{
    "scenarios": {
        "getobject": {
            "seq": [
                "authenticate",
                "get-object-by-id"
            ]
        }
    },
    "requests": {
        "authenticate": {
            "method": "POST",
            "url": "https://${oauthbaseurl}/sso/oauth2/token",
            "headers": {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            "body": "${secret}",
            "output": [
                {
                    "key": "token",
                    "value": "$.access_token"
                }
            ]
        },
        "get-object-by-id": {
            "method": "GET",
            "url": "https://${baseurl}/objects/${objectid}",
            "headers": {
                "Authorization": "Bearer ${token}"
            }
        }
    }
}
```

`-e dev` refers to an env file with defaults (`dev.env.json`):

```
{
    "inputs": [
        {
            "key": "oauthbaseurl",
            "value": "example.com"
        },
        {
            "key": "secret",
            "value": "SECRET"
        },
        {
            "key": "baseurl",
            "value": "example.com"
        }
    ]
}
```

`objectid=12345` provides an extra input, merged with defaults