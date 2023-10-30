package util

import (
	"fmt"
	"time"
)

// MessageURL formats URL link for message from text channel of guild with guild ID, channel ID, message ID
func MessageURL(guildID, channelID, messageID string) string {
	return fmt.Sprintf("https://discord.com/channels/%s/%s/%s", guildID, channelID, messageID)
}

// FormatKoreanTime formats time to time in Korea (UTC +09:00)
func FormatKoreanTime(time time.Time) string {
	// TODO: Implement
	return ""
}
