# Backend Return

## Json Struct

```
{
   "status": bool, // if success. When it's true, just read "data"
   "action": str // the action witch hope the frontend executes when "status" is false
   "data": any // data by request
}
```

## Action Meaning

### redirect

The frontend need switch the page by "data", maybe the page is a new url, or just update page rendering.

```
"data": {
   "url": str,
   "args": [
      ...
   ]
}
```

### alert

The frontend need display the msg in "data".

```
"data": str // msg
```

# Middleware

 - `cors`: allow cors, just using in dev
 - `session`: Session by `github.com/gin-contrib/sessions`
 - `csrf`: Csrf by `github.com/utrack/gin-csrf`
 - `needLogin`: check if login by session status

# Session

status list:

 - `ifLogined`: bool
    - `true`: already logined
    - `false`: not login