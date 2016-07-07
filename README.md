# Micha

[![Build Status](https://travis-ci.org/onrik/micha.svg?branch=master)](https://travis-ci.org/onrik/micha)
[![Coverage Status](https://coveralls.io/repos/github/onrik/micha/badge.svg?branch=master)](https://coveralls.io/github/onrik/micha?branch=master)
[![Gitter](https://badges.gitter.im/onrik/micha.svg)](https://gitter.im/onrik/micha)

Telegram bot framework

- [x] Sending messages
- [x] Sending files
- [x] Forward messages
- [x] Updating messages
- [x] Inline mode
- [ ] Inline mode cached results 
- [ ] Getting files
- [ ] Answer callback query
- [ ] Getting updates by webhook


##### Simple echo bot example:
```go
package main

import (
	"github.com/onrik/micha"
	"log"
)

func main() {
	bot, err := micha.NewBot("<token>")
	if err != nil {
		log.Fatal(err)
	}

	go bot.Start()

	for update := range bot.Updates {
		if update.Message != nil {
			bot.SendMessage(update.Message.Chat.Id, update.Message.Text, nil)
		}
	}
}

```
