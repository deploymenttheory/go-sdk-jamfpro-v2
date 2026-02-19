package client

import (
	"crypto/tls"
	"fmt"
	"maps"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// ClientOption configures the Transport.
type ClientOption func(*Transport) error

// WithBaseURL sets a custom base URL for the API client.
func WithBaseURL(baseURL string) ClientOption {
	return func(t *Transport) error {
		t.BaseURL = baseURL
		t.logger.Info("Base URL configured", zap.String("base_url", baseURL))
		return nil
	}
}

// WithTimeout sets a custom timeout for HTTP requests.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetTimeout(timeout)
		t.logger.Info("HTTP timeout configured", zap.Duration("timeout", timeout))
		return nil
	}
}

// WithRetryCount sets the number of retries for failed requests.
func WithRetryCount(count int) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryCount(count)
		t.logger.Info("Retry count configured", zap.Int("retry_count", count))
		return nil
	}
}

// WithRetryWaitTime sets the wait time between retry attempts.
func WithRetryWaitTime(waitTime time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryWaitTime(waitTime)
		t.logger.Info("Retry wait time configured", zap.Duration("wait_time", waitTime))
		return nil
	}
}

// WithRetryMaxWaitTime sets the maximum wait time between retries.
func WithRetryMaxWaitTime(maxWaitTime time.Duration) ClientOption {
	return func(t *Transport) error {
		t.client.SetRetryMaxWaitTime(maxWaitTime)
		t.logger.Info("Retry max wait time configured", zap.Duration("max_wait_time", maxWaitTime))
		return nil
	}
}

// WithLogger sets a custom logger for the client.
func WithLogger(logger *zap.Logger) ClientOption {
	return func(t *Transport) error {
		if logger == nil {
			return fmt.Errorf("logger cannot be nil")
		}
		t.logger = logger
		t.logger.Info("Custom logger configured")
		return nil
	}
}

// WithDebug enables debug mode which logs request and response details.
func WithDebug() ClientOption {
	return func(t *Transport) error {
		t.client.SetDebug(true)
		t.logger.Info("Debug mode enabled")
		return nil
	}
}

// WithUserAgent sets a custom user agent string.
func WithUserAgent(userAgent string) ClientOption {
	return func(t *Transport) error {
		t.client.SetHeader("User-Agent", userAgent)
		t.userAgent = userAgent
		t.logger.Info("User agent configured", zap.String("user_agent", userAgent))
		return nil
	}
}

// WithGlobalHeader sets a global header included in all requests.
func WithGlobalHeader(key, value string) ClientOption {
	return func(t *Transport) error {
		t.globalHeaders[key] = value
		t.logger.Info("Global header configured", zap.String("key", key), zap.String("value", value))
		return nil
	}
}

// WithGlobalHeaders sets multiple global headers at once.
func WithGlobalHeaders(headers map[string]string) ClientOption {
	return func(t *Transport) error {
		maps.Copy(t.globalHeaders, headers)
		t.logger.Info("Multiple global headers configured", zap.Int("count", len(headers)))
		return nil
	}
}

// WithProxy sets an HTTP proxy for all requests.
func WithProxy(proxyURL string) ClientOption {
	return func(t *Transport) error {
		t.client.SetProxy(proxyURL)
		t.logger.Info("Proxy configured", zap.String("proxy", proxyURL))
		return nil
	}
}

// WithTLSClientConfig sets custom TLS configuration.
func WithTLSClientConfig(tlsConfig *tls.Config) ClientOption {
	return func(t *Transport) error {
		t.client.SetTLSClientConfig(tlsConfig)
		t.logger.Info("TLS client config configured")
		return nil
	}
}

// WithTransport sets a custom HTTP transport (http.RoundTripper).
func WithTransport(transport http.RoundTripper) ClientOption {
	return func(t *Transport) error {
		t.client.SetTransport(transport)
		t.logger.Info("Custom transport configured")
		return nil
	}
}

// WithInsecureSkipVerify disables TLS certificate verification (use only for testing).
func WithInsecureSkipVerify() ClientOption {
	return func(t *Transport) error {
		t.client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
		t.logger.Warn("TLS certificate verification DISABLED - use only for testing")
		return nil
	}
}

// WithMaxConcurrentRequests sets the maximum number of concurrent API requests.
// Jamf Pro guidance recommends no more than 5 concurrent connections to avoid
// disrupting other Jamf Pro tasks and managed devices. Pass 0 to disable.
func WithMaxConcurrentRequests(n int) ClientOption {
	return func(t *Transport) error {
		if n <= 0 {
			t.sem = nil
			t.logger.Info("Concurrency limiting disabled")
			return nil
		}
		t.sem = newSemaphore(n)
		t.logger.Info("Concurrency limit configured", zap.Int("max_concurrent_requests", n))
		return nil
	}
}

// WithMandatoryRequestDelay sets a fixed delay after every successful request.
// Use for bulk operations to avoid hitting Jamf Pro rate limits.
func WithMandatoryRequestDelay(d time.Duration) ClientOption {
	return func(t *Transport) error {
		t.requestDelay = d
		t.logger.Info("Mandatory request delay configured", zap.Duration("delay", d))
		return nil
	}
}

// WithTotalRetryDuration sets a maximum total wall-clock budget for a request
// including all retry attempts. Requests that exceed this duration are cancelled.
func WithTotalRetryDuration(d time.Duration) ClientOption {
	return func(t *Transport) error {
		t.totalRetryDuration = d
		t.logger.Info("Total retry duration configured", zap.Duration("total_retry_duration", d))
		return nil
	}
}
