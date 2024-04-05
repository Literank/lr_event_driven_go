/*
Package remote provides remote service communication funcs.
*/
package remote

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const defaultTimeout = 5 * time.Second

// HttpFetch fetches resources from a http URL.
func HttpFetch(ctx context.Context, url string, result interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Timeout: defaultTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, result)
	if err != nil {
		return err
	}
	return nil
}
