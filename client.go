package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Client ...
type Client struct {
	Client       *http.Client
	baseURL      *url.URL
	apiKey       string
	username     string
	password     string
	clientID     string
	clientSecret string
	token        *token
	log          *zap.SugaredLogger
}

// NewClient create new Ghost client.
func NewClient(baseURL, apiKey, username, password, clientID, clientSecret string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed parsing base url")
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, errors.Wrap(err, "failed creating logger")
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Client{
		Client:       http.DefaultClient,
		baseURL:      u,
		apiKey:       apiKey,
		username:     username,
		password:     password,
		clientID:     clientID,
		clientSecret: clientSecret,
		log:          sugar.Named("go-ghost"),
	}, nil
}

// Post fetches post from Ghost using post ID.
func (c *Client) Post(id string) (*Post, error) {
	u := *c.baseURL
	u.Path = path.Join(u.Path, "ghost", "api", "v2", "content", "posts", id)
	q := u.Query()
	q.Add("key", c.apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed creating request")
	}
	req.Header.Set("Accept", "application/json")

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed making request ")
	}
	defer res.Body.Close()

	if !validStatusCode(res.StatusCode) {
		return nil, fmt.Errorf("invalid status code: %s", res.Status)
	}

	data := Payload{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "failed decoding response")
	}

	if len(data.Errors) > 0 {
		return nil, fmt.Errorf("message: %s context: %s", data.Errors[0].Message, data.Errors[0].Context)
	}

	return &data.Posts[0], nil
}

// CreatePost creates new post.
func (c *Client) CreatePost(post *Post) (*Post, error) {
	if c.token == nil {
		err := c.auth()
		if err != nil {
			return nil, errors.Wrap(err, "failed authentication")
		}
	}

	u := *c.baseURL
	u.Path = path.Join(u.Path, "ghost", "api", "v0.1", "posts")

	b := &bytes.Buffer{}
	br := Payload{
		Posts: []Post{*post},
	}
	err := json.NewEncoder(b).Encode(br)

	req, err := http.NewRequest(http.MethodPost, u.String(), b)
	if err != nil {
		return nil, errors.Wrap(err, "failed creating request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token.accessToken)

	c.log.Debugf("auth: querying %s", u.String())
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed making request ")
	}
	defer res.Body.Close()

	if !validStatusCode(res.StatusCode) {
		return nil, fmt.Errorf("invalid status code: %s", res.Status)
	}

	data := Payload{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, errors.Wrap(err, "failed decoding response")
	}

	if len(data.Errors) > 0 {
		return nil, fmt.Errorf("message: %s context: %s", data.Errors[0].Message, data.Errors[0].Context)
	}

	return &data.Posts[0], nil
}

func (c *Client) auth() error {
	u := *c.baseURL
	u.Path = path.Join(u.Path, "ghost", "api", "v0.1", "authentication", "token")

	v := url.Values{}
	v.Set("grant_type", "password")
	v.Set("username", c.username)
	v.Set("password", c.password)
	v.Set("client_id", c.clientID)
	v.Set("client_secret", c.clientSecret)
	c.log.Debug("auth: credentials: ", v.Encode())
	b := strings.NewReader(v.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), b)
	if err != nil {
		return errors.Wrap(err, "failed creating request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	c.log.Debugf("auth: querying %s", u.String())
	res, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed making request")
	}
	defer res.Body.Close()

	if !validStatusCode(res.StatusCode) {
		return fmt.Errorf("invalid status code: %s", res.Status)
	}

	data := Payload{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return errors.Wrap(err, "failed decoding response")
	}

	if len(data.Errors) > 0 {
		return fmt.Errorf("message: %s context: %s", data.Errors[0].Message, data.Errors[0].Context)
	}

	c.token = &token{
		accessToken:  data.AccessToken,
		refreshToken: data.RefreshToken,
		expiration:   time.Now().Add(time.Duration(data.ExpiresIn) * time.Second),
	}
	c.log.Debugf("auth: access token: %s refresh token: %s expiration: %s", c.token.accessToken, c.token.refreshToken, c.token.expiration.String())

	return nil
}

type token struct {
	accessToken  string
	refreshToken string
	expiration   time.Time
}

func validStatusCode(statusCode int) bool {
	if statusCode < http.StatusOK || statusCode >= http.StatusInternalServerError {
		return false
	}
	if statusCode >= http.StatusMultipleChoices && statusCode <= http.StatusBadRequest {
		return false
	}
	return true
}
