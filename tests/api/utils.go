/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/apecloud/kb-cloud-client-go/tests"

	"golang.org/x/net/publicsuffix"
)

// WithFakeAuth avoids issue of API returning `text/html` instead of `application/json`
func WithFakeAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		common.ContextAPIKeys,
		map[string]common.APIKey{
			"apiKeyAuth": {
				Key: "00000000000000000000000000000000",
			},
			"appKeyAuth": {
				Key: "0000000000000000000000000000000000000000",
			},
		},
	)
}

// WithTestAuth returns authenticated context.
func WithTestAuth(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		common.ContextDigestAuth,
		common.DigestAuth{
			UserName: os.Getenv("KB_CLOUD_TEST_API_KEY_NAME"),
			Password: os.Getenv("KB_CLOUD_TEST_API_KEY_SECRET"),
		},
	)
}

// NewDefaultContext return context with detected values.
func NewDefaultContext(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	if site, ok := os.LookupEnv("KB_CLOUD_TEST_SITE"); ok {
		ctx = context.WithValue(
			ctx,
			common.ContextServerIndex,
			2,
		)
		ctx = context.WithValue(
			ctx,
			common.ContextServerVariables,
			map[string]string{"site": site},
		)
	}
	return ctx
}

// NewConfiguration return configuration with known options.
func NewConfiguration() *common.Configuration {
	config := common.NewConfiguration()
	config.Debug = os.Getenv("DEBUG") == "true"
	config.RetryConfiguration.EnableRetry = true
	return config
}

type contextKey string

var (
	clientKey = contextKey("client")
)

// WithClient sets client for unit tests in context.
func WithClient(ctx context.Context) context.Context {
	ctx = NewDefaultContext(ctx)
	return context.WithValue(ctx, clientKey, common.NewAPIClient(NewConfiguration()))
}

// ClientFromContext returns client and indication if it was successful.
func ClientFromContext(ctx context.Context) (*common.APIClient, bool) {
	if ctx == nil {
		return nil, false
	}
	v := ctx.Value(clientKey)
	if c, ok := v.(*common.APIClient); ok {
		return c, true
	}
	return nil, false
}

// Client returns client from context.
func Client(ctx context.Context) *common.APIClient {
	c, ok := ClientFromContext(ctx)
	if !ok {
		log.Fatal("client is not configured")
	}
	return c
}

// WithRecorder configures client with recorder.
func WithRecorder(ctx context.Context, t *testing.T) (context.Context, func()) {
	ctx = WithClient(ctx)
	client := Client(ctx)

	ctx, err := tests.WithClock(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup clock: %v", err)
	}

	r, err := tests.Recorder(ctx, tests.SecurePath(t.Name()))
	if err != nil {
		t.Fatalf("could not setup recorder: %v", err)
	}
	client.GetConfig().HTTPClient = &http.Client{Transport: tests.WrapRoundTripper(r)}

	return ctx, func() {
		r.Stop()
	}
}

func getTestDomain(ctx context.Context, client *common.APIClient) (string, error) {
	baseURL, err := client.GetConfig().ServerURLWithContext(ctx, "")
	if err != nil {
		return "", fmt.Errorf("could not generate base url: %v", err)
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("could not parse base url: %v", err)
	}

	host, err := publicsuffix.EffectiveTLDPlusOne(parsedURL.Host)
	if err != nil {
		return "", fmt.Errorf("could not parse TLD+1: %v", err)
	}
	return host, nil
}

// SendRequest sends request to endpoints without specification.
func SendRequest(ctx context.Context, method, url string, payload []byte) (*http.Response, []byte, error) {
	baseURL := ""
	if !strings.HasPrefix(url, "https://") {
		var err error
		baseURL, err = Client(ctx).GetConfig().ServerURLWithContext(ctx, "")
		if err != nil {
			return nil, []byte{}, fmt.Errorf("failed to get base URL for Datadog API: %s", err.Error())
		}
	}

	request, err := http.NewRequest(method, baseURL+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("failed to create request for Datadog API: %s", err.Error())
	}
	keys := ctx.Value(common.ContextAPIKeys).(map[string]common.APIKey)
	request.Header.Add("DD-API-KEY", keys["apiKeyAuth"].Key)
	request.Header.Add("DD-APPLICATION-KEY", keys["appKeyAuth"].Key)
	request.Header.Set("Content-Type", "application/json")

	resp, respErr := Client(ctx).GetConfig().HTTPClient.Do(request)
	body, rerr := io.ReadAll(resp.Body)
	if rerr != nil {
		respErr = fmt.Errorf("failed reading response body: %s", rerr.Error())
	}
	return resp, body, respErr
}
