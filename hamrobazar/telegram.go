package hamrobazar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func msgTelegram(apiToken, chatID, msg, imgUrl string) error {
	data := map[string]interface{}{
		"chat_id":    chatID,
		"caption":    msg,
		"parse_mode": "HTML",
		"photo":      imgUrl,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", apiToken)
	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", string(body))
	}

	return nil
}
