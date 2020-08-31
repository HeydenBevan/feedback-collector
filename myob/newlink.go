package myob

// NewLink is the object that defines a new link between Domain settings and Forum Output
type NewLink struct {
	Domain       string        `json:"domain"`
	ForumLabel   string        `json:"forumlabel"`
	ForumTag     string        `json:"forumtag"`
	Destinations []Destination `json:"destinations"`
}

// Destination object contains the names for which destinations the processor will use
type Destination struct {
	Name string `json:"name"`
}
