# GO-SLACK

[![Documentation](https://godoc.org/github.com/faruqisan/go-slack?status.svg)](https://godoc.org/github.com/faruqisan/go-slack) [![Go Report Card](https://goreportcard.com/badge/github.com/faruqisan/go-slack)](https://goreportcard.com/report/github.com/faruqisan/go-slack)

## A simple golang slack API

Support Multiple Webhooks

simple because it only send given message to a webhook url(s)


## Example

### setup

- **import the package**

    ```go
    import "github.com/faruqisan/go-slack"
    ```
- **setup option**

    ```go
        opt := slack.Option{
            WebHookURLs: ["PUT_YOUR_WEBHOOK_URL_HERE"],
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

    //or Send Async
    sl.SendAsync("message")
    ```
## TODO

- create unit test
