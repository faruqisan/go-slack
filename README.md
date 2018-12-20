# GO-SLACK

[![Documentation](https://godoc.org/github.com/faruqisan/go-slack?status.svg)](https://godoc.org/github.com/faruqisan/go-slack) [![Go Report Card](https://goreportcard.com/badge/github.com/faruqisan/go-slack)](https://goreportcard.com/report/github.com/faruqisan/go-slack)

## A simple golang slack API

simple because it only send given message to a webhook url

## Example

### setup

- **import the package**

    ```go
    import "github.com/faruqisan/go-slack"
    ```
- **setup option**

    ```go
        opt := slack.Option{
            WebHookURL: "PUT_YOUR_WEBHOOK_URL_HERE",
        }
    ```
- **create the object with option**

    ```go
        sl := slack.New(opt)
    ```

### send the message

- **use send function**
    ```go
    err := sl.Send("your message here")
    if err != nil {
        //handle the error
    }
    ```
## TODO

- create unit test
