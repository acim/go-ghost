package ghost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
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
}

// NewClient create new Ghost client.
func NewClient(baseURL, apiKey, username, password, clientID, clientSecret string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed parsing base url")
	}
	return &Client{
		Client:       http.DefaultClient,
		baseURL:      u,
		apiKey:       apiKey,
		username:     username,
		password:     password,
		clientID:     clientID,
		clientSecret: clientSecret,
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
	req.Header.Add("Accept", "application/json")

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed making request ")
	}
	defer res.Body.Close()

	data := PostsResponse{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}

	for _, post := range data.Posts {
		fmt.Printf("%#v\n", post)
	}
	for _, err := range data.Errors {
		fmt.Printf("%#v\n", err)
	}
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
	b := strings.NewReader(v.Encode())

	req, err := http.NewRequest(http.MethodPost, u.String(), b)
	if err != nil {
		return errors.Wrap(err, "failed creating request")
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := c.Client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed making request ")
	}
	defer res.Body.Close()

	data := authResponse{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return errors.Wrap(err, "failed decoding response")
	}

	c.token = &token{
		accessToken:  data.AccessToken,
		refreshToken: data.RefreshToken,
		expiration:   time.Now().Add(time.Duration(data.ExpiresIn) * time.Second),
	}

	return nil
}

type token struct {
	accessToken  string
	refreshToken string
	expiration   time.Time
}
