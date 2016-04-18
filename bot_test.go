package micha

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BotTestSuite struct {
	suite.Suite
	bot *Bot
}

func (s *BotTestSuite) SetupSuite() {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.telegram.org/bot111/getMe",
		httpmock.NewStringResponder(200, `{"ok":true,"result":{"id":1,"first_name":"Micha","username":"michabot"}}`))

	bot, err := NewBot("111")
	s.Equal(err, nil)
	s.Equal(bot.Me.FirstName, "Micha")
	s.Equal(bot.Me.Id, int64(1))
	s.Equal(bot.Me.Username, "michabot")

	s.bot = bot
}

func (s *BotTestSuite) TestBuildUrl() {
	url := s.bot.buildUrl("someMethod")
	s.Equal(url, "https://api.telegram.org/bot111/someMethod")
}

func TestBotTestSuite(t *testing.T) {
	suite.Run(t, new(BotTestSuite))
}
