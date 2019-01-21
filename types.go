package ghost

import (
	"time"
)

// Payload ...
type Payload struct {
	Posts        []Post  `json:"posts,omitempty"`
	Users        []User  `json:"users,omitempty"`
	Tags         []Tag   `json:"tags,omitempty"`
	AccessToken  string  `json:"access_token,omitempty"`
	RefreshToken string  `json:"refresh_token,omitempty"`
	ExpiresIn    int64   `json:"expires_in,omitempty"`
	TokenType    string  `json:"token_type,omitempty"`
	Meta         Meta    `json:"meta,omitempty"`
	Errors       []Error `json:"errors,omitempty"`
}

// Post ...
type Post struct {
	ID                 string    `json:"id,omitempty"`
	UUID               string    `json:"uuid,omitempty"`
	Title              string    `json:"title,omitempty"`
	Slug               string    `json:"slug,omitempty"`
	Mobiledoc          string    `json:"mobiledoc,omitempty"` // https://github.com/bustle/mobiledoc-kit/blob/master/MOBILEDOC.md
	HTML               string    `json:"html,omitempty"`
	CommendID          string    `json:"comment_id,omitempty"`
	PlaintText         string    `json:"plaintext,omitempty"`
	FeatureImage       string    `json:"feature_image,omitempty"`
	Featured           bool      `json:"featured,omitempty"`
	Page               bool      `json:"page,omitempty"`
	Status             string    `json:"status,omitempty"`
	Locale             string    `json:"locale,omitempty"`
	Visibility         string    `json:"visibility,omitempty"`
	MetaTitle          string    `json:"meta_title,omitempty"`
	MetaDescription    string    `json:"meta_description,omitempty"`
	AuthorID           string    `json:"author_id,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	CreatedBy          string    `json:"created_by,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	UpdatedBy          string    `json:"updated_by,omitempty"`
	PublishedAt        time.Time `json:"published_at,omitempty"`
	PublishedBy        string    `json:"published_by,omitempty"`
	CustomExcerpt      string    `json:"custom_excerpt,omitempty"`
	CodeinjectionHead  string    `json:"codeinjection_head,omitempty"`
	CodeinjectionFoot  string    `json:"codeinjection_foot,omitempty"`
	OgImage            string    `json:"og_image,omitempty"`
	OgTitle            string    `json:"og_title,omitempty"`
	OgDescription      string    `json:"og_description,omitempty"`
	TwitterImage       string    `json:"twitter_image,omitempty"`
	TwitterTitle       string    `json:"twitter_title,omitempty"`
	TwitterDescription string    `json:"twitter_description,omitempty"`
	CustomTemplate     string    `json:"custom_template,omitempty"`
}

// User ...
type User struct {
	ID                   string    `json:"id,omitempty"`
	Name                 string    `json:"Name,omitempty"`
	Slug                 string    `json:"slug,omitempty"`
	GhostAuthAccessToken string    `json:"ghost_auth_access_token,omitempty"`
	GhostAuthID          string    `json:"ghost_auth_id,omitempty"`
	Password             string    `json:"password,omitempty"`
	Email                string    `json:"email,omitempty"`
	ProfileImage         string    `json:"profile_image,omitempty"`
	CoverImage           string    `json:"cover_image,omitempty"`
	Bio                  string    `json:"bio,omitempty"`
	Website              string    `json:"website,omitempty"`
	Location             string    `json:"location,omitempty"`
	Facebook             string    `json:"facebook,omitempty"`
	Twitter              string    `json:"twitter,omitempty"`
	Accessibility        string    `json:"accessibility,omitempty"`
	Status               string    `json:"status,omitempty"`
	Locale               string    `json:"locale,omitempty"`
	Visibility           string    `json:"visibility,omitempty"`
	MetaTitle            string    `json:"meta_title,omitempty"`
	MetaDescription      string    `json:"meta_description,omitempty"`
	Tour                 string    `json:"tour,omitempty"`
	LastSeen             time.Time `json:"last_seen,omitempty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	CreatedBy            string    `json:"created_by,omitempty"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
	UpdatedBy            string    `json:"updated_by,omitempty"`
}

// Tag ...
type Tag struct {
	ID              string    `json:"id,omitempty"`
	Name            string    `json:"Name,omitempty"`
	Slug            string    `json:"slug,omitempty"`
	Description     string    `json:"description,omitempty"`
	FeatureImage    string    `json:"feature_image,omitempty"`
	ParentID        string    `json:"parent_id,omitempty"`
	Visibility      string    `json:"visibility,omitempty"`
	MetaTitle       string    `json:"meta_title,omitempty"`
	MetaDescription string    `json:"meta_description,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	CreatedBy       string    `json:"created_by,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	UpdatedBy       string    `json:"updated_by,omitempty"`
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

// Error ...
type Error struct {
	Message   string `json:"message,omitempty"`
	Context   string `json:"context,omitempty"`
	ErrorType string `json:"errorType,omitempty"`
}

// Mobiledoc ...
type Mobiledoc struct {
	Version  string          `json:"version"`
	Atoms    []interface{}   `json:"atoms"`
	Cards    [][]interface{} `json:"cards"`
	Markups  []interface{}   `json:"markups"`
	Sections [][]interface{} `json:"sections"`
}

// Card ...
type Card struct {
	HTML     *string `json:"html,omitempty"`
	Markdown *string `json:"markdown,omitempty"`
}
