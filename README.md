# sendgrid-go
A wrapper around sendgrid/sendgrid-go library to consume the api easily

# install the package
run `go get github.com/vikashvikram/sendgrid-go`

# example
Following example will send mail to list of participant specified in `to` with subject `Test Mail` and plain text content `This is a test mail`.

```go
package main

import (
	"fmt"
	"log"

	"github.com/vikashvikram/sendgrid-go"
)

func main() {
	from := "no-reply@example.com"
	to := []string{"email1@example.com", "email2@example.com"}
	subject := "Test Mail"
	body := "This is a test mail"
	response, err := sendgrid.Mail(from, subject, to, body)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```
### Please Note
1. please make sure that SENDGRID_API_KEY env variable exists with valid values. In case you do not have it, please generate one from sendgrid.com website
2. Sometimes you receive mails in your junk box so please don't forget to check mails there. In such cases, you will need to tell your email service that this is not a junk mail, so that in future you receive mails sent from sendgrid in your inbox.

