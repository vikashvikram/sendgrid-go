# sendgrid-go
A wrapper around sendgrid/sendgrid-go library to consume the api easily

# install the package
run `go get github.com/vikashvikram/sendgrid-go`

# example
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
	subject := "testing sendgrid module"
	body := "Ignore it. I am just testing"
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
### please make sure that SENDGRID_API_KEY env variable exists with valid values. In case you do not have it, please generate one from sendgrid.com website

