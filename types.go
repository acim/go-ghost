package ghost

// PostsReqRes ...
type PostsReqRes struct {
	Posts  []Post  `json:"posts,omitempty"`
	Meta   Meta    `json:"meta,omitempty"`
	Errors []Error `json:"errors,omitempty"`
}

// Meta ...
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

// Pagination ...
type Pagination struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Pages int64 `json:"pages"`
	Total int64 `json:"total"`
	Next  int64 `json:"next"`
	Prev  int64 `json:"prev"`
}

// Post ...
type Post struct {
	ID                 string `json:"id,omitempty"`
	UUID               string `json:"uuid,omitempty"`
	Title              string `json:"title,omitempty"`
	Slug               string `json:"slug,omitempty"`
	HTML               string `json:"html,omitempty"`
	CommentID          string `json:"comment_id,omitempty"`
	FeatureImage       string `json:"feature_image,omitempty"`
	Featured           bool   `json:"featured,omitempty"`
	Page               bool   `json:"page,omitempty"`
	Status             string `json:"status,omitempty"`
	Locale             string `json:"locale,omitempty"`
	Visibility         string `json:"visibility,omitempty"`
	MetaTitle          string `json:"meta_title,omitempty"`
	MetaDescription    string `json:"meta_description"`
	CreatedAt          string `json:"created_at,omitempty"`
	CreatedBy          string `json:"created_by,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	UpdatedBy          string `json:"updated_by,omitempty"`
	PublishedAt        string `json:"published_at,omitempty"`
	PublishedBy        string `json:"published_by,omitempty"`
	CustomExcerpt      string `json:"custom_excerpt,omitempty"`
	CodeinjectionHead  string `json:"codeinjection_head,omitempty"`
	CodeinjectionFoot  string `json:"codeinjection_foot,omitempty"`
	OgImage            string `json:"og_image,omitempty"`
	OgTitle            string `json:"og_title,omitempty"`
	OgDescription      string `json:"og_description,omitempty"`
	TwitterImage       string `json:"twitter_image,omitempty"`
	TwitterTitle       string `json:"twitter_title,omitempty"`
	TwitterDescription string `json:"twitter_description,omitempty"`
	CustomTemplate     string `json:"custom_template,omitempty"`
	PrimaryAuthor      string `json:"primary_author,omitempty"`
	PrimaryTag         string `json:"primary_tag,omitempty"`
	URL                string `json:"url,omitempty"`
	Excerpt            string `json:"excerpt,omitempty"`
}

// Error ...
type Error struct {
	Message   string `json:"message"`
	Context   string `json:"context"`
	ErrorType string `json:"errorType"`
}

type authRes struct {
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	ExpiresIn    int64   `json:"expires_in"`
	TokenType    string  `json:"token_type"`
	Errors       []Error `json:"errors"`
}
