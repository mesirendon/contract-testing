//go:build contracts

package users

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/pact-foundation/pact-go/v2/consumer"
	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/matchers"
	"github.com/stretchr/testify/assert"
)

func TestUserClientPact_GetUser(t *testing.T) {
	_ = log.SetLogLevel("INFO")

	mockProvider, err := consumer.NewV2Pact(consumer.MockHTTPProviderConfig{
		Consumer: os.Getenv("CONSUMER_NAME"),
		Provider: os.Getenv("PROVIDER_NAME"),
		LogDir:   os.Getenv("LOG_DIR"),
		PactDir:  os.Getenv("PACT_DIR"),
	})

	t.Run("the user exists", func(t *testing.T) {
		id := 10

		err = mockProvider.
			AddInteraction().
			Given("User drwho exists").
			UponReceiving("A request to login with user 'drwho'").
			WithRequestPathMatcher("GET", matchers.Regex("/user/"+strconv.Itoa(id), "/user/[0-9]+"), func(vrb *consumer.V2RequestBuilder) {
				vrb.Header("Authorization", matchers.Like("Bearer 2016-01-01T05:43"))
			}).
			WillRespondWith(http.StatusOK, func(vrb *consumer.V2ResponseBuilder) {
				vrb.BodyMatch(user{}).
					Header("Content-Type", matchers.Term("application/json", `application\/json`)).
					Header("X-Api-Correlation-Id", matchers.Like("100"))
			}).
			ExecuteTest(t, func(msc consumer.MockServerConfig) error {
				u := fmt.Sprintf("http://%s:%d", msc.Host, msc.Port)
				client, _ := NewUsersClient(u)
				client.SetToken("2016-01-01T05:43")

				user, err := client.GetUser(id)
				if user.ID != id {
					return fmt.Errorf("wanted user with ID %d but got %d", id, user.ID)
				}

				return err
			})

		assert.NoError(t, err)
	})

	t.Run("the user does not exist", func(t *testing.T) {
		id := 10

		err = mockProvider.
			AddInteraction().
			Given("User drwho does not exist").
			UponReceiving("A request to login with user 'drwho'").
			WithRequestPathMatcher("GET", matchers.Regex("/user/"+strconv.Itoa(id), "/user/[0-9]+"), func(vrb *consumer.V2RequestBuilder) {
				vrb.Header("Authorization", matchers.Like("Bearer 2016-01-01T05:43"))
			}).
			WillRespondWith(http.StatusNotFound, func(vrb *consumer.V2ResponseBuilder) {
				vrb.
					Header("Content-Type", matchers.Term("application/json", `application\/json`)).
					Header("X-Api-Correlation-Id", matchers.Like("100"))
			}).
			ExecuteTest(t, func(msc consumer.MockServerConfig) error {
				u := fmt.Sprintf("http://%s:%d", msc.Host, msc.Port)
				client, _ := NewUsersClient(u)
				client.SetToken("2016-01-01T05:43")

				_, err := client.GetUser(id)
				assert.Equal(t, ErrNotFound, err)

				return nil
			})

		assert.NoError(t, err)
	})

	t.Run("the user is not authenticated", func(t *testing.T) {
		id := 10

		err = mockProvider.
			AddInteraction().
			Given("User is not authenticated").
			UponReceiving("A request to get a user").
			WithRequestPathMatcher("GET", matchers.Regex("/user/"+strconv.Itoa(id), "/user/[0-9]+")).
			WillRespondWith(http.StatusUnauthorized, func(vrb *consumer.V2ResponseBuilder) {
				vrb.
					Header("Content-Type", matchers.Term("application/json", `application\/json`)).
					Header("X-Api-Correlation-Id", matchers.Like("100"))
			}).
			ExecuteTest(t, func(msc consumer.MockServerConfig) error {
				u := fmt.Sprintf("http://%s:%d", msc.Host, msc.Port)
				client, _ := NewUsersClient(u)

				_, err := client.GetUser(id)
				assert.Equal(t, ErrUnauthorized, err)

				return nil
			})

		assert.NoError(t, err)
	})
}
