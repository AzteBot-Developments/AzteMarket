package logUtils

import (
	"github.com/RazvanBerbece/AzteMarket/pkg/config"
	"github.com/RazvanBerbece/AzteMarket/src/libs/models/events"
	loggerService "github.com/RazvanBerbece/AzteMarket/src/libs/services/logger/strategies"
	"github.com/bwmarrin/discordgo"
)

func PublishConsoleLogErrorEvent(logsChannel chan events.LogEvent, msg string) {
	logsChannel <- events.LogEvent{
		Logger: loggerService.NewConsoleLogger(),
		Msg:    msg,
		Type:   "ERROR",
	}
}

func PublishDiscordLogErrorEvent(logsChannel chan events.LogEvent, s *discordgo.Session, topicName string, channelsConfig []string, msg string) {
	logsChannel <- events.LogEvent{
		Logger: loggerService.NewDiscordChannelLogger(
			s,
			config.GetDiscordChannelIdForTopic(topicName, channelsConfig),
			true,
		),
		Msg:  msg,
		Type: "ERROR",
	}
}

func PublishDiscordLogWarnEvent(logsChannel chan events.LogEvent, s *discordgo.Session, topicName string, channelsConfig []string, msg string) {
	logsChannel <- events.LogEvent{
		Logger: loggerService.NewDiscordChannelLogger(
			s,
			config.GetDiscordChannelIdForTopic(topicName, channelsConfig),
			true,
		),
		Msg:  msg,
		Type: "WARN",
	}
}

func PublishConsoleLogInfoEvent(logsChannel chan events.LogEvent, msg string) {
	logsChannel <- events.LogEvent{
		Logger: loggerService.NewConsoleLogger(),
		Msg:    msg,
		Type:   "INFO",
	}
}

func PublishDiscordLogInfoEvent(logsChannel chan events.LogEvent, s *discordgo.Session, topicName string, channelsConfig []string, msg string) {
	logsChannel <- events.LogEvent{
		Logger: loggerService.NewDiscordChannelLogger(
			s,
			config.GetDiscordChannelIdForTopic(topicName, channelsConfig),
			true,
		),
		Msg:  msg,
		Type: "INFO",
	}
}
