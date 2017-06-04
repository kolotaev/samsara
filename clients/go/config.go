package client

import (
	"time"
)

// Config is Samsara SDK default configuration.
type Config struct {
	// Samsara ingestion api endpoint "http://samsara-ingestion.local/"
	url string

	// Identifier of the source of these events.
	// OPTIONAL used only for record-event
	sourceId string

	// Start the publishing thread?
	// default = true
	startPublishingThread bool

	// How often should the events being sent to Samsara
	// in milliseconds.
	// default = 30s
	publishInterval uint32

	// Max size of the buffer.
	// When buffer is full older events are dropped.
	maxBufferSize int64

	// Minimum number of events that must be in the buffer
	// before attempting to publish them.
	minBufferSize int64

	// Network timeout for send operations
	// in milliseconds.
	// default 30s
	sendTimeout uint32

	// Should the payload be compressed?
	// allowed values :gzip, :none
	compression string

	// NOT CURRENTLY SUPPORTED
	// Add Samsara client statistics events
	// this helps you to understand whether the
	// buffer size and publish-intervals are
	// adequately configured.
	// sendClientStats bool
}

// NewConfig creates a new Config with all default values.
// Later you may want to replace them with specific values.
func NewConfig() Config {
	config := Config{}
	config.url = ""
	config.sourceId = ""
	config.startPublishingThread = true
	config.publishInterval = 30000
	config.maxBufferSize = 10000
	config.minBufferSize = 100
	config.sendTimeout = 30000
	config.compression = "gzip"
	//config.sendClientStats = true
	return config
}

// Validate validates given configuration values.
func (c *Config) Validate() error {
	switch {
	case len(c.url) == 0:
		return ConfigValidationError{"URL for Ingestion API should be specified."}
	case c.compression != "gzip" && c.compression != "none":
		return ConfigValidationError{"Incorrect compression option."}
	case c.publishInterval <= 0:
		return ConfigValidationError{"Invalid interval time for Samsara client."}
	case c.maxBufferSize < c.minBufferSize:
		return ConfigValidationError{"maxBufferSize can not be less than minBufferSize."}
	default:
		return nil
	}
}

// Timestamp generates current timestamp.
func Timestamp() int64 {
	return time.Now().UnixNano() / 1000000
}
