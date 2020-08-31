package myob

// FeedbackCollector defines a family of methods that return Feedback Collectors
type FeedbackCollector interface {
	DomainName() string
	SeeFeedback() *FeedbackContent
}

// Feedback captures the Forum feedback coming in.
type Feedback struct {
	Source          string            `json:"source"`
	FeedbackContent []FeedbackContent `json:"feedback-content"`
}

// FeedbackContent is a universal Wrapper for individual Feedback Posts
type FeedbackContent struct {
	Author       string
	Title        string
	ContentBody  string
	ExternalLink string
	DatePosted   string
}

// NewFeedback initialises a FeedbackContent wrapper
func NewFeedback() *FeedbackContent {
	var fc FeedbackContent
	return &fc
}
