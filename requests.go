package tasty

import (
	"fmt"
	"io"
	"net/http"
)

func doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Authorization", token)

	if req.Method == "POST" {
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request errored with status = %d: %s", resp.StatusCode, respBody)
	}

	return respBody, nil
}
