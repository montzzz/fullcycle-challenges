package utils

import (
	"fullcycle-auction_go/configuration/logger"
	"os"
	"time"
)

func GetAuctionDuration() time.Duration {
	auctionInterval := os.Getenv("AUCTION_INTERVAL")
	if auctionInterval == "" {
		logger.Info("AUCTION_INTERVAL not set, defaulting to 5 minutes")
		return time.Minute * 5
	}

	duration, err := time.ParseDuration(auctionInterval)
	if err != nil {
		logger.Error("AUCTION_INTERVAL not set, defaulting to 5 minutes", err)
		return time.Minute * 5
	}

	logger.Info("Auction interval set to " + auctionInterval)

	return duration
}
