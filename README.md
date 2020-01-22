# UniOne

[![Travis](https://img.shields.io/travis/alexeyco/unione.svg)](https://travis-ci.org/alexeyco/unione)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/unione/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/unione?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/unione)](https://goreportcard.com/report/github.com/alexeyco/unione)
[![GoDoc](https://godoc.org/github.com/alexeyco/unione?status.svg)](https://godoc.org/github.com/alexeyco/unione)

Send transactional emails with [one.unisender.com](https://one.unisender.com).
Examples and documentation available on [godoc](https://godoc.org/github.com/alexeyco/unione).

# Usage
## Send email

```go
package main

import (
    "log"

    "github.com/alexeyco/unione"
    "github.com/alexeyco/unione/message"
)

func main() {
	recipient := message.NewRecipient("recipient@example.com").
		Name("John Doe")

	msg := message.NewMessage().
		From("site@example.com", "My site").
		To(recipient).
		Subject("Awesome news, buddy").
		BodyPlainText("Return to my site and enjoy")

	client := unione.New("username", "api-key")

	if err := client.Send(msg); err != nil {
		log.Fatalln(err)
	}
}
```

## License
```
MIT License

Copyright (c) 2019 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
