package service

const WhatsAppAPIURL = "https://graph.facebook.com/v17.0/<YOUR_PHONE_NUMBER_ID>/messages"
const AccessToken = "<YOUR_WHATSAPP_API_ACCESS_TOKEN>"

type WhatsAppMessage struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

func SendWhatsAppMessage(to, message string) error {
	//msg := WhatsAppMessage{
	//	MessagingProduct: "whatsapp",
	//	To:               to,
	//	Type:             "text",
	//}
	//msg.Text.Body = message
	//
	//body, err := json.Marshal(msg)
	//if err != nil {
	//	return err
	//}
	//
	//req, err := http.NewRequest("POST", WhatsAppAPIURL, bytes.NewBuffer(body))
	//if err != nil {
	//	return err
	//}
	//
	//req.Header.Set("Authorization", "Bearer "+AccessToken)
	//req.Header.Set("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != http.StatusOK {
	//	return errors.New("failed to send WhatsApp message")
	//}

	return nil
}
