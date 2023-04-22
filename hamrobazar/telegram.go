package hamrobazar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func msgTelegram(msg, imgUrl string) error {
	tgToken := os.Getenv("HAMROBAZAR_TG_TOKEN")
	tgChatID := os.Getenv("HAMROBAZAR_TG_CHAT_ID")
	url := fmt.Sprintf("https://api.telegram.org/%s/sendPhoto", tgToken)

	data := map[string]interface{}{
		"chat_id":    tgChatID,
		"caption":    msg,
		"parse_mode": "HTML",
		"photo":      imgUrl,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", string(body))
	}

	return nil
}
