# httperr

httperr is a simple interface for easily implementing errors in http responses.

It's as simple as calling JSON on a basic error.

```go
    httperr.JSON(w,r,
        &httperr.DefaultError{
                Message:    "There was an error retrieving the items",
                ErrorCode:  "GI0001",
                StatusCode: 500,
        },
    )
```

The default `DefaultError` changes the status code of the response on `Render`, then is marshalled.

This was inspired by [render](https://github.com/go-chi/render), but simplified and modified to fit the usual simple error handling
