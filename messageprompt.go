package main

type messagePrompt struct {
	messageQueue []string
}

func (m *messagePrompt) push(message string) {
	m.messageQueue = append(m.messageQueue, message)
}

func (m *messagePrompt) pop() string {
	message := m.messageQueue[0]
	messages.messageQueue = messages.messageQueue[1:]
	return message
}

func (m messagePrompt) processTurn() {

}

func (m messagePrompt) processKey(char rune) (validKey bool) {

	switch char {
	case 0:
		return true
	default:
		return false
	}
}
