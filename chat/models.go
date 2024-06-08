package chat

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages []Message `json:"messages"`
	Model    string    `json:"model"`
}

type ChatResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message      Message `json:"message"`
	LogProbs     float32 `json:"log_probs"`
	FinishReason string  `json:"finish_reason"`
	Index        int     `json:"index"`
}
