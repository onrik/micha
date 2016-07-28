package micha

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type BotTestSuite struct {
	suite.Suite
	bot *Bot
}

func (s *BotTestSuite) SetupSuite() {
	httpmock.Activate()

	httpmock.RegisterResponder("GET", "https://api.telegram.org/bot111/getMe",
		httpmock.NewStringResponder(200, `{"ok":true,"result":{"id":1,"first_name":"Micha","username":"michabot"}}`))

	bot, err := NewBot("111")
	s.Equal(err, nil)
	s.Equal(bot.Me.FirstName, "Micha")
	s.Equal(bot.Me.Id, int64(1))
	s.Equal(bot.Me.Username, "michabot")

	s.bot = bot
}

func (s *BotTestSuite) TearDownSuite() {
	httpmock.DeactivateAndReset()
}

func (s *BotTestSuite) registerResponse(method string, params url.Values, response string) {
	url := s.bot.buildUrl(method)
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, response))
}

func (s *BotTestSuite) TestBuildUrl() {
	url := s.bot.buildUrl("someMethod")
	s.Equal(url, "https://api.telegram.org/bot111/someMethod")
}

func (s *BotTestSuite) TestGetChat() {
	s.registerResponse("getChat", url.Values{"chat_id": {"123"}}, `{
		"ok": true,
		"result": {
			"id": 123,
			"type": "group",
			"title": "ChatTitle",
			"first_name": "fn",
			"last_name": "ln",
			"username": "un"
		}
	}`)

	chat, err := s.bot.GetChat(123)
	s.Equal(err, nil)
	s.Equal(chat.Id, int64(123))
	s.Equal(chat.Type, CHAT_TYPE_GROUP)
	s.Equal(chat.Title, "ChatTitle")
	s.Equal(chat.FirstName, "fn")
	s.Equal(chat.LastName, "ln")
	s.Equal(chat.Username, "un")
}

func (s *BotTestSuite) TestGetChatAdministrators() {
	s.registerResponse("getChatAdministrators", url.Values{"chat_id": {"123"}}, `{
		"ok": true,
		"result": [
			{
				"status": "administrator",
				"user": {
					"id": 456,
					"first_name": "John",
					"last_name": "Doe",
					"username": "john_doe"
				}
			},
			{
				"status": "administrator",
				"user": {
					"id": 789,
					"first_name": "Mohammad",
					"last_name": "Li",
					"username": "mli"
				}
			}
		]
	}`)

	administrators, err := s.bot.GetChatAdministrators(123)
	s.Equal(err, nil)
	s.Equal(len(administrators), 2)
	s.Equal(administrators[0].User.Id, int64(456))
	s.Equal(administrators[0].User.FirstName, "John")
	s.Equal(administrators[0].User.LastName, "Doe")
	s.Equal(administrators[0].User.Username, "john_doe")
	s.Equal(administrators[0].Status, MEMBER_STATUS_ADMINISTRATOR)

	s.Equal(administrators[1].User.Id, int64(789))
	s.Equal(administrators[1].User.FirstName, "Mohammad")
	s.Equal(administrators[1].User.LastName, "Li")
	s.Equal(administrators[1].User.Username, "mli")
	s.Equal(administrators[1].Status, MEMBER_STATUS_ADMINISTRATOR)
}

func (s *BotTestSuite) TestGetChatMember() {
	s.registerResponse("getChatMember", url.Values{"chat_id": {"123"}, "user_id": {"456"}}, `{
		"ok": true,
		"result": {
			"status": "creator",
			"user": {
				"id": 456,
				"first_name": "John",
				"last_name": "Doe",
				"username": "john_doe"
			}
		}
	}`)

	chatMember, err := s.bot.GetChatMember(123, 456)
	s.Equal(err, nil)
	s.Equal(chatMember.User.Id, int64(456))
	s.Equal(chatMember.User.FirstName, "John")
	s.Equal(chatMember.User.LastName, "Doe")
	s.Equal(chatMember.User.Username, "john_doe")
	s.Equal(chatMember.Status, MEMBER_STATUS_CREATOR)

}

func (s *BotTestSuite) TestGetChatMembersCount() {
	s.registerResponse("getChatMembersCount", url.Values{"chat_id": {"123"}}, `{"ok":true, "result": 25}`)

	count, err := s.bot.GetChatMembersCount(123)
	s.Equal(err, nil)
	s.Equal(count, 25)
}

func (s *BotTestSuite) TestGetFile() {
	s.registerResponse("getFile", url.Values{"file_id": {"222"}}, `{"ok":true,"result":{"file_id":"222","file_size":5,"file_path":"document/file_3.txt"}}`)

	file, err := s.bot.GetFile("222")
	s.Equal(err, nil)
	s.Equal(file.FileId, "222")
	s.Equal(file.FileSize, uint64(5))
	s.Equal(file.FilePath, "document/file_3.txt")
}

func TestBotTestSuite(t *testing.T) {
	suite.Run(t, new(BotTestSuite))
}
