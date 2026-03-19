package micha

// AnswerPreCheckoutQuery
// Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query.
// Use this method to respond to such pre-checkout queries. On success, True is returned.
// Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
func (bot *Bot) AnswerPreCheckoutQuery(queryID string, ok bool, errorMessage string) error {
	params := map[string]any{
		"pre_checkout_query_id": queryID,
		"ok":                    ok,
	}
	if errorMessage != "" {
		params["error_message"] = errorMessage
	}

	err := bot.Post("answerPreCheckoutQuery", params, nil)

	return err
}
