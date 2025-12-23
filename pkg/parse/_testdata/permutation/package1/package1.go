package package1

// +Foo=true
// +Bar=123
type AnotherUser struct {
	// +ID=true
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}
