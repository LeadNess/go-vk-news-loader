package api

type VKWall struct {
	Count int64    `json:"count"`
	Items []VKPost `json:"items"`
}

type VKPost struct {
	ID          int64           `json:"id"`
	Date        int64           `json:"date"`
	PostType    string          `json:"post_type"`
	Text        string          `json:"text"`
	IsPinned    int8            `json:"is_pinned"`
	Comments    struct {
		Count		int64			 `json:"count"`
	}                           `json:"comments"`
	Likes       struct {
		Count		int64			 `json:"count"`
	}                           `json:"likes"`
	Reposts     struct {
		Count		int64			 `json:"count"`
	}                           `json:"reposts"`
	Views       struct {
		Count		int64			 `json:"count"`
	}                           `json:"views"`
	Attachments []VKAttachments `json:"attachments"`
}

type VKAttachments struct {
	Type  string   `json:"type"`
	Link  struct {
		Url          string   `json:"url"`
		Title        string   `json:"title"`
		Description  string   `json:"description"`
	}              `json:"link"`
}