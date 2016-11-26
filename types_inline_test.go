package micha

import (
	"testing"
)

func TestInlineQueryResult(t *testing.T) {
	(InlineQueryResultArticle{}).itsInlineQueryResult()
	(InlineQueryResultPhoto{}).itsInlineQueryResult()
	(InlineQueryResultCachedPhoto{}).itsInlineQueryResult()
	(InlineQueryResultGif{}).itsInlineQueryResult()
	(InlineQueryResultCachedGif{}).itsInlineQueryResult()
	(InlineQueryResultMpeg4Gif{}).itsInlineQueryResult()
	(InlineQueryResultCachedMpeg4Gif{}).itsInlineQueryResult()
	(InlineQueryResultVideo{}).itsInlineQueryResult()
	(InlineQueryResultCachedVideo{}).itsInlineQueryResult()
	(InlineQueryResultAudio{}).itsInlineQueryResult()
	(InlineQueryResultCachedAudio{}).itsInlineQueryResult()
	(InlineQueryResultVoice{}).itsInlineQueryResult()
	(InlineQueryResultCachedVoice{}).itsInlineQueryResult()
	(InlineQueryResultDocument{}).itsInlineQueryResult()
	(InlineQueryResultCachedDocument{}).itsInlineQueryResult()
	(InlineQueryResultLocation{}).itsInlineQueryResult()
	(InlineQueryResultVenue{}).itsInlineQueryResult()
	(InlineQueryResultCachedSticker{}).itsInlineQueryResult()
	(InlineQueryResultContact{}).itsInlineQueryResult()
	(InlineQueryResultGame{}).itsInlineQueryResult()
}

func TestInputMessageContent(t *testing.T) {
	(InputTextMessageContent{}).itsInputMessageContent()
	(InputLocationMessageContent{}).itsInputMessageContent()
	(InputVenueMessageContent{}).itsInputMessageContent()
	(inputMessageContentImplementation{}).itsInputMessageContent()
}
