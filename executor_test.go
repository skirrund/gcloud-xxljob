package gcloudxxljob

import (
	"context"
	"errors"
	"testing"

	"github.com/skirrund/gcloud/logger"
)

func TestXxx(t *testing.T) {
	opts := Options{
		AdminAddresses:   "http://xxljob:8089/xxl-job-admin",
		AppName:          "xxl-job-test-go",
		Logretentiondays: 1,
	}
	executor := Init(opts)
	executor.RegTask("go-test", JobRun)
	executor.Run()
}
func JobRun(ctx context.Context, req *RunRequest) error {
	logger.Info("run job")
	return errors.New("test")
}
