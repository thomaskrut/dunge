package main

type messagePrompt struct {
	messageQueue []string
}

func (m *messagePrompt) addMessage(message string) {
	m.messageQueue = append(m.messageQueue, message)
}

func (m *messagePrompt) getOldestMessage() string {
	return m.messageQueue[0]
}

func (m *messagePrompt) deleteOldestMessage() {
	messages.messageQueue = messages.messageQueue[1:]
}

func (m messagePrompt) processTurn() {
	
}

func (m messagePrompt) processKey(char rune) bool {
	
	switch char {
		case 0: {
			
			return true
		}
	}
	return false
}
