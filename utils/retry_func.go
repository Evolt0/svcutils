package utils

import (
	"context"
	"time"
)

func RetryFunc(function func() error, retryCount int64, retryDurationMillSecond time.Duration) (err error) {
	timeOut := time.NewTicker(retryDurationMillSecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(retryCount)*retryDurationMillSecond)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return err
		case <-timeOut.C:
			err = function()
			if err == nil {
				return nil
			}
		}
	}
}
