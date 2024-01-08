package hamrobazar

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type hamrobazar struct {
	telegramAPIToken string
	telegramChatID   string
}

func NewHamrobazar(telegramAPIToken string, telegramChatID string) *hamrobazar {
	return &hamrobazar{
		telegramAPIToken: telegramAPIToken,
		telegramChatID:   telegramChatID,
	}
}

// Run executes a filter query against a database and sends messages via Telegram
func (t *hamrobazar) Run(filter Filter) error {
	records, err := readDatabase()
	if err != nil {
		return fmt.Errorf("failed to read database: %w", err)
	}

	items, err := getData(filter)
	if err != nil {
		return fmt.Errorf("failed to get data: %w", err)
	}
	fmt.Printf("Found %d records.\n", len(items))

	for _, item := range items {
		if contains(records, item.ID) {
			continue
		}

		message := formatMessage(item.Name, item.Price, item.Location.LocationDescription, item.CreatorInfo, item.Description, getItemURL(item.ID))
		imgURL := formatImgURL(item.ImageURL)
		if err := msgTelegram(t.telegramAPIToken, t.telegramChatID, message, imgURL); err != nil {
			return fmt.Errorf("failed to send message: %w", err)
		}

		records = append(records, item.ID)
	}

	if err := storeToDatabase(records); err != nil {
		return fmt.Errorf("failed to store to database: %w", err)
	}

	return nil
}

func getData(filter Filter) ([]Item, error) {
	req, err := http.NewRequest("GET", "https://api.hamrobazaar.com/api/Product", nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = filter.FilterParams.URLQuery()

	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("Apikey", "09BECB8F84BCB7A1796AB12B98C1FB9E") // Appears to be a public API key
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Deviceid", "87a1f113-fa6b-4501-9255-e242411d47ab") // Random UUID
	req.Header.Set("Devicesource", "web")
	req.Header.Set("Referer", "https://hamrobazaar.com/")
	req.Header.Set("Referrer-Policy", "strict-origin-when-cross-origin")
	req.Header.Set("Strict-Transport-Security", "max-age=2592000")
	req.Header.Set("X-Content-Type-Options", "nosniff")
	req.Header.Set("X-Frame-Options", "SAMEORIGIN")

	var client http.Client
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result SearchResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	if !result.Succeeded {
		return nil, errors.New(result.Message)
	}

	return result.Data, nil
}
