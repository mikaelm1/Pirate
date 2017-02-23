package data

// Subscription model
type Subscription struct {
	Subscribed bool   `json:"subscribed"`
	CreatedAt  string `json:"created_at"`
	RepoURL    string `json:"repository_url"`
}
