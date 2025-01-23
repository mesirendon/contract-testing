package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/mesirendon/contract-testing/consumer/internal/model"
)

var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrUnavailable  = errors.New("api unavailable")
)

type UsersClient struct {
	baseURL    *url.URL
	httpClient *http.Client
	token      string
}

func NewUsersClient(baseURL string) (*UsersClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parsing base url: %w", err)
	}

	return &UsersClient{
		baseURL:    u,
		httpClient: http.DefaultClient,
	}, nil
}

func (uc *UsersClient) SetToken(token string) {
	uc.token = token
}

func (uc *UsersClient) GetUser(id int) (model.User, error) {
	req, err := uc.newRequest("GET", fmt.Sprintf("/user/%d", id), nil)
	if err != nil {
		return model.User{}, err
	}
	var usr user
	res, err := uc.do(req, &usr)
	if res != nil {
		switch res.StatusCode {
		case http.StatusNotFound:
			return model.User{}, ErrNotFound
		case http.StatusUnauthorized:
			return model.User{}, ErrUnauthorized
		case http.StatusInternalServerError:
			return model.User{}, ErrUnavailable
		}
	}

	if err != nil {
		return model.User{}, ErrUnavailable
	}

	return usr.toModel(), err
}

func (uc *UsersClient) do(req *http.Request, v any) (*http.Response, error) {
	resp, err := uc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

func (uc *UsersClient) newRequest(method, path string, body any) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := uc.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if uc.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", uc.token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Admin Service")

	return req, nil
}
