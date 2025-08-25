package dify

import (
	"context"
	"net/http"
)

type ChatMessageRequest struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	ResponseMode   string                 `json:"response_mode"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	User           string                 `json:"user"`
}

type ChatMessageResponse1 struct {
	ID             string `json:"id"`
	Answer         string `json:"answer"`
	ConversationID string `json:"conversation_id"`
	CreatedAt      int    `json:"created_at"`
}


type ChatMessageResponse struct {
    Answer         string    `json:"answer"`
    ConversationID string    `json:"conversation_id"`
    CreatedAt      int64     `json:"created_at"`
    Event          string    `json:"event"`
    ID             string    `json:"id"`
    MessageID      string    `json:"message_id"`
    Metadata       *Metadata `json:"metadata"`
    Mode           string    `json:"mode"`
    TaskID         string    `json:"task_id"`
}

type Metadata struct {
    AnnotationReply     interface{}   `json:"annotation_reply"`
    RetrieverResources  []interface{} `json:"retriever_resources"`
    Usage               *Usage        `json:"usage"`
}

type Usage struct {
    CompletionPrice      string `json:"completion_price"`
    CompletionPriceUnit  string `json:"completion_price_unit"`
    CompletionTokens     int    `json:"completion_tokens"`
    CompletionUnitPrice  string `json:"completion_unit_price"`
    Currency             string `json:"currency"`
    Latency              int    `json:"latency"`
    PromptPrice          string `json:"prompt_price"`
    PromptPriceUnit      string `json:"prompt_price_unit"`
    PromptTokens         int    `json:"prompt_tokens"`
    PromptUnitPrice      string `json:"prompt_unit_price"`
    TotalPrice           string `json:"total_price"`
    TotalTokens          int    `json:"total_tokens"`
}

/* Create chat message
 * Create a new conversation message or continue an existing dialogue.
 */
func (api *API) ChatMessages(ctx context.Context, req *ChatMessageRequest) (resp *ChatMessageResponse, err error) {
	req.ResponseMode = "blocking"

	httpReq, err := api.createBaseRequest(ctx, http.MethodPost, "/v1/chat-messages", req)
	if err != nil {
		return
	}
	err = api.c.sendJSONRequest(httpReq, &resp)
	return
}
