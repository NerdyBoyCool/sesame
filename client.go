package sesame

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/chmike/cmac-go"
)

const (
	baseURL = "https://app.candyhouse.co/api/sesame2"
)

// Client is a client of sesame API.
type Client struct {
	HTTPClient *http.Client
	APIKey     string
	SecretKey  string
	DeviseUUID string
	URL        string
}

// APIError represents an error of connpass API.
type APIError struct {
	StatusCode int
}

// Error implements error.Error.
func (err *APIError) Error() string {
	return fmt.Sprintf("StatusCode: %d", err.StatusCode)
}

// BodyParams is a Request Body for API Request
type BodyParams struct {
	Cmd     int16  `json:"cmd"`
	History string `json:"history"`
	Sign    string `json:"sign"`
}

// NewClient creates a new sesame api client.
func NewClient(apiKey, secretKey, deviseUUID string) *Client {
	var cli Client
	cli = Client{
		HTTPClient: http.DefaultClient,
		APIKey:     apiKey,
		SecretKey:  secretKey,
		DeviseUUID: deviseUUID,
		URL:        baseURL,
	}
	return &cli
}

func (client *Client) signature() (string, error) {
	i := int32(time.Now().Unix())
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	byteKey, err := hex.DecodeString(client.SecretKey)
	if err != nil {
		return "", err
	}
	cm, err := cmac.New(aes.NewCipher, byteKey)
	if err != nil {
		return "", err
	}
	cm.Write(buf[1:4])
	m := cm.Sum(nil)
	return hex.EncodeToString(m), nil
}

func (client *Client) do(ctx context.Context, cmd int16, history string) (*http.Response, error) {
	sign, err := client.signature()
	if err != nil {
		return nil, fmt.Errorf("cannot generate signature: %w", err)
	}
	from := base64.StdEncoding.EncodeToString([]byte(history))
	data, err := json.Marshal(&BodyParams{Cmd: cmd, Sign: sign, History: from})
	if err != nil {
		return nil, fmt.Errorf("cannot marshal to json: %w", err)
	}
	req, err := http.NewRequest("POST", "https://app.candyhouse.co/api/sesame2/"+client.DeviseUUID+"/cmd", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("cannot create HTTP request: %w", err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("x-api-key", client.APIKey)
	return client.HTTPClient.Do(req)
}

// Lock Sesame
func (client *Client) Lock(ctx context.Context, history string) error {
	resp, err := client.do(ctx, 82, history)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return &APIError{StatusCode: resp.StatusCode}
	}
	return nil
}

// Unlock Sesame
func (client *Client) Unlock(ctx context.Context, history string) error {
	resp, err := client.do(ctx, 83, history)
	defer resp.Body.Close()

	if err != nil {
		return err
	}

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return &APIError{StatusCode: resp.StatusCode}
	}

	return nil
}

// Toggle Sesame
func (client *Client) Toggle(ctx context.Context, history string) error {
	resp, err := client.do(ctx, 88, history)
	defer resp.Body.Close()

	if err != nil {
		return fmt.Errorf("cannot toogle sesame: %w", err)
	}

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return &APIError{StatusCode: resp.StatusCode}
	}

	return nil
}
