package config

import (
	"strings"
)

func GetDiscordChannelIdForTopic(topic string, topicPairs []string) string {

	for _, channelPairString := range topicPairs {
		channelValues := strings.Split(channelPairString, "-")
		topicName := channelValues[0]
		if topicName == topic {
			return channelValues[1]
		}
	}

	return ""

}
