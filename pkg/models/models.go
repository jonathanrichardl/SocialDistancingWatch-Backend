package models

type MqttAlert struct {
	Classroom       string `json:"class"`
	TotalViolations string `json:"distance"`
}

type ViolationData struct {
	Time            string `json:"time"`
	TotalViolations string `json:"violations"`
	Class           string `json:"class"`
}

type TelegramOutgoingMessage struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

type TelegramIncomingMessage struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}