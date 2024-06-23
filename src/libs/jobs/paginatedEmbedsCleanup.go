package backgroundJobs

import (
	"fmt"
	"time"

	actionEventsUtils "github.com/RazvanBerbece/AzteMarket/src/handlers/remoteEvents/actionEvents/utils"
	sharedRuntime "github.com/RazvanBerbece/AzteMarket/src/shared/runtime"
	"github.com/bwmarrin/discordgo"
)

func ClearOldPaginatedEmbeds(s *discordgo.Session, sFrequency int, tThreshold time.Duration) {

	fmt.Println("[CRON] Starting Task ClearOldPaginatedEmbeds() at", time.Now(), "running every", sFrequency, "seconds")

	ticker := time.NewTicker(time.Duration(sFrequency) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				go cleanupOldPaginatedEmbeds(s, tThreshold)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

}

func cleanupOldPaginatedEmbeds(s *discordgo.Session, threshold time.Duration) {
	for msgId, embedData := range sharedRuntime.EmbedsToPaginate {
		if time.Since(time.Unix(int64(embedData.Timestamp), 0)) > threshold {
			go actionEventsUtils.DisablePaginatedEmbed(s, embedData.ChannelId, msgId)
			delete(sharedRuntime.EmbedsToPaginate, msgId)
		}
	}
}
