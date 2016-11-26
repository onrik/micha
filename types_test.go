package micha

import (
	"testing"
)

func TestReplyMarkup(t *testing.T) {
	(ReplyKeyboardMarkup{}).itsReplyMarkup()
	(ReplyKeyboardRemove{}).itsReplyMarkup()
	(InlineKeyboardMarkup{}).itsReplyMarkup()
	(ForceReply{}).itsReplyMarkup()
	(ForceReply{}).itsReplyMarkup()
}
