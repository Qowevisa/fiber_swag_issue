# How to reproduce the issue:

1. Launch main.go:
```
go run main.go
```

2. Go to `http://localhost:8888/api/swagger`

3. Click on the `Authorize` button

4. Enter `user` as Username and click `Authorize`. Then click `Close`

5. Try to execute `/auth/ping` endpoint

## How to "fix" it

1. Instead of immediately clicking on `Close` after Username was set to `user` go to `Password` and put something there and delete it. That way it is forced to be an empty string
