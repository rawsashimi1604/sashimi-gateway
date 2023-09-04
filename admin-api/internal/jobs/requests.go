package jobs

import "github.com/rawsashimi1604/sashimi-gateway/admin-api/internal/api/analytics"

// TODO: create gateway for adding requests.
type RequestCronJob struct {
	AnalyticsTracker *analytics.AnalyticsTracker
}
