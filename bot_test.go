package micha

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	httpmock.Deactivate()
}

func (s *BotTestSuite) TearDownTest() {
	httpmock.Reset()
}

func (s *BotTestSuite) registerResponse(method string, params url.Values, response string) {
	url := s.bot.buildUrl(method)
	if params != nil {
		url += fmt.Sprintf("?%s", params.Encode())
	}

	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, response))
}

func (s *BotTestSuite) registerRequestCheck(method string, exceptedRequest string) {
	url := s.bot.buildUrl(method)

	httpmock.RegisterResponder("POST", url, func(request *http.Request) (*http.Response, error) {
		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}

		s.Equal(exceptedRequest, strings.TrimSpace(string(body)))
		return httpmock.NewStringResponse(200, `{"ok":true, "result": {}}`), nil
	})
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

func (s *BotTestSuite) TestDownloadFileUrl() {
	url := s.bot.DownloadFileUrl("file.mp3")
	s.Equal(url, "https://api.telegram.org/file/bot111/file.mp3")
}

func (s *BotTestSuite) TestSendPhoto() {
	request := `{"chat_id":111,"photo":"35f9f497a879436fbb6e682f6dd75986","caption":"test caption","reply_to_message_id":143}`
	s.registerRequestCheck("sendPhoto", request)

	message, err := s.bot.SendPhoto(111, "35f9f497a879436fbb6e682f6dd75986", &SendPhotoOptions{
		Caption:          "test caption",
		ReplyToMessageId: 143,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendAudio() {
	request := `{"chat_id":123,"audio":"061c2810391f44f6beffa3ee8a7e5af4","duration":36,"performer":"John Doe","title":"Single","reply_to_message_id":143}`
	s.registerRequestCheck("sendAudio", request)

	message, err := s.bot.SendAudio(123, "061c2810391f44f6beffa3ee8a7e5af4", &SendAudioOptions{
		Duration:         36,
		Performer:        "John Doe",
		Title:            "Single",
		ReplyToMessageId: 143,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendDocument() {
	request := `{"chat_id":124,"document":"efd8d08958894a6781873b9830634483","caption":"document caption","reply_to_message_id":144}`
	s.registerRequestCheck("sendDocument", request)

	message, err := s.bot.SendDocument(124, "efd8d08958894a6781873b9830634483", &SendDocumentOptions{
		Caption:          "document caption",
		ReplyToMessageId: 144,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendSticker() {
	request := `{"chat_id":125,"sticker":"070114a7fa964322acb3d65e6e36eb2b","reply_to_message_id":145}`
	s.registerRequestCheck("sendSticker", request)

	message, err := s.bot.SendSticker(125, "070114a7fa964322acb3d65e6e36eb2b", &SendStickerOptions{
		ReplyToMessageId: 145,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendVideo() {
	request := `{"chat_id":126,"video":"b169f647c020405b8c9035cf3f315ff0","duration":22,"width":320,"height":240,"caption":"video caption","reply_to_message_id":146}`
	s.registerRequestCheck("sendVideo", request)

	message, err := s.bot.SendVideo(126, "b169f647c020405b8c9035cf3f315ff0", &SendVideoOptions{
		Duration:         22,
		Width:            320,
		Height:           240,
		Caption:          "video caption",
		ReplyToMessageId: 146,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendVoice() {
	request := `{"chat_id":127,"voice":"75ac50947bc34a3ea2efdca5000d9ad5","duration":56,"reply_to_message_id":147}`
	s.registerRequestCheck("sendVoice", request)

	message, err := s.bot.SendVoice(127, "75ac50947bc34a3ea2efdca5000d9ad5", &SendVoiceOptions{
		Duration:         56,
		ReplyToMessageId: 147,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendLocation() {
	request := `{"chat_id":128,"latitude":22.532434,"longitude":-44.8243324,"reply_to_message_id":148}`
	s.registerRequestCheck("sendLocation", request)

	message, err := s.bot.SendLocation(128, 22.532434, -44.8243324, &SendLocationOptions{
		ReplyToMessageId: 148,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendVenue() {
	request := `{"chat_id":129,"latitude":22.532434,"longitude":-44.8243324,"title":"Kremlin","address":"Red Square 1","foursquare_id":"1","reply_to_message_id":149}`
	s.registerRequestCheck("sendVenue", request)

	message, err := s.bot.SendVenue(129, 22.532434, -44.8243324, "Kremlin", "Red Square 1", &SendVenueOptions{
		FoursquareId:     "1",
		ReplyToMessageId: 149,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendContact() {
	request := `{"chat_id":130,"phone_number":"+79998887766","first_name":"John","last_name":"Doe","reply_to_message_id":150}`
	s.registerRequestCheck("sendContact", request)

	message, err := s.bot.SendContact(130, "+79998887766", "John", "Doe", &SendContactOptions{
		ReplyToMessageId: 150,
	})

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestForwardMessage() {
	request := `{"chat_id":131,"disable_notification":true,"from_chat_id":99,"message_id":543}`
	s.registerRequestCheck("forwardMessage", request)

	message, err := s.bot.ForwardMessage(131, 99, 543, true)

	s.Equal(err, nil)
	s.NotEqual(message, nil)
}

func (s *BotTestSuite) TestSendChatAction() {
	request := `{"action":"typing","chat_id":132}`
	s.registerRequestCheck("sendChatAction", request)

	err := s.bot.SendChatAction(132, CHAT_ACTION_TYPING)
	s.Equal(err, nil)
}

func TestBotTestSuite(t *testing.T) {
	suite.Run(t, new(BotTestSuite))
}
