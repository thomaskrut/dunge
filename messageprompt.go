package main

type messagePrompt struct {
	messageQueue []string
}

func (m *messagePrompt) addMessage(message string) {
	m.messageQueue = append(m.messageQueue, message)
}

func (m messagePrompt) processTurn() {
}

func (m messagePrompt) processKey(char rune) {
}
