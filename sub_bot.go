package telebot

func (b *Bot) SubBot(token string) *Bot {
	return &Bot{
		Token:       token,
		URL:         b.URL,
		onError:     b.onError,
		group:       b.group,
		handlers:    b.handlers,
		synchronous: b.synchronous,
		verbose:     b.verbose,
		parseMode:   b.parseMode,
		client:      b.client,
	}
}

func (b *Bot) GetMe() (*User, error) {
	if b.Me == nil {
		return b.getMe()
	}

	return b.Me, nil
}
