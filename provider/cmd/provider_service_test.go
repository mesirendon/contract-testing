//go:build contracts

package main

import (
	"fmt"
	l "log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/mesirendon/contract-testing/provider/internal/middleware"
	"github.com/pact-foundation/pact-go/v2/log"
	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/pact-foundation/pact-go/v2/utils"
)

const (
	timeFormat = "2006-01-02T15:04"
)

var (
	port, _ = utils.GetFreePort()
)

func TestUserServicePact(t *testing.T) {
	_ = log.SetLogLevel("INFO")

	go startInstrumentedUserService()

	err := provider.NewVerifier().VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		BrokerURL:                  fmt.Sprintf("%s://%s", os.Getenv("PACT_BROKER_PROTO"), os.Getenv("PACT_BROKER_URL")),
		ProviderBranch:             os.Getenv("VERSION_BRANCH"),
		Provider:                   os.Getenv("PROVIDER_NAME"),
		BrokerUsername:             os.Getenv("PACT_BROKER_USERNAME"),
		BrokerPassword:             os.Getenv("PACT_BROKER_PASSWORD"),
		FailIfNoPactsFound:         false,
		PublishVerificationResults: true,
		ProviderVersion:            os.Getenv("VERSION_COMMIT"),
		RequestFilter:              fixBearerToken,
	})

	if err != nil {
		t.Log(err)
	}
}

func fixBearerToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", time.Now().Format(timeFormat)))
		}
		next.ServeHTTP(w, r)
	})
}

func startInstrumentedUserService() {
	mux := middleware.GetHTTPHandler()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		l.Fatal(err)
	}
	defer ln.Close()

	l.Printf("API starting: port %d (%s)", port, ln.Addr())
	l.Printf("API terminating: %v", http.Serve(ln, mux))
}
