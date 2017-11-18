package resp

// Action type is client-side behaviour type which depends on API response.
type ActionType int
const (
	ActionTypeNone = iota
	ActionTypeReload
	ActionTypeGotoTop
	ActionTypeDelegateToClient
)

// Result of API call.
type APIResult struct {
	Code 		int `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ActionType  ActionType  `json:"action_type"`
	Data        interface{} `json:"data"`
}