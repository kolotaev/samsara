package client

import (
	"strconv"
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	expected := Config{
		url:                   "",
		sourceId:              "",
		startPublishingThread: true,
		publishInterval:       30000,
		maxBufferSize:         10000,
		minBufferSize:         100,
		sendTimeout:           30000,
		compression:           "gzip",
	}

	initial := NewConfig()
	if expected != initial {
		t.Error("New config doesn't contain correct default values")
	}
}

func TestConfig_Validate_WithCorrectData(t *testing.T) {
	config := NewConfig()

	config.startPublishingThread = false
	config.sourceId = "some-source-id"
	config.publishInterval = 20
	config.sendTimeout = 10
	config.compression = "none"
	config.url = "http://foo.bar"
	config.maxBufferSize = 8
	config.minBufferSize = 4

	if err := config.Validate(); err != nil {
		t.Errorf("Config with correct data should not raise errors. Given config = %+v, Error = %s", config, err)
	}
}

func TestConfig_Validate_WithInvalidData(t *testing.T) {
	sets := []struct {
		msg    string
		config Config
	}{
		{
			"URL for Ingestion API should be specified.",
			func() Config {
				config := NewConfig()
				config.url = ""
				return config
			}(),
		},
		{
			"Invalid interval time for Samsara client.",
			func() Config {
				config := NewConfig()
				config.url = "http://foo.bar"
				config.publishInterval = 0
				return config
			}(),
		},
		{
			"Incorrect compression option.",
			func() Config {
				config := NewConfig()
				config.url = "http://foo.bar"
				config.compression = "zip"
				return config
			}(),
		},
		{
			"maxBufferSize can not be less than minBufferSize.",
			func() Config {
				config := NewConfig()
				config.url = "http://foo.bar"
				config.maxBufferSize = 4
				config.minBufferSize = 8
				return config
			}(),
		},
	}
	for i, v := range sets {
		err := v.config.Validate()
		if err == nil {
			t.Errorf("Set #%d. Expected error validation for Config %+v", i, v.config)
		}
		if err.Error() != v.msg {
			t.Errorf("Set #%d. Expected error message %q, got %q", i, v.msg, err.Error())
		}
	}
}

func TestConfig_Timestamp(t *testing.T) {
	now := Timestamp()
	time.Sleep(1 * time.Millisecond)
	later := Timestamp()

	if now >= later {
		t.Errorf("Real ticking timestamp should be returned. Expected %d > %d", later, now)
	}

	if len(strconv.Itoa(int(now))) != 13 {
		t.Errorf("Timestamp should be in milliseconds. Got %d", now)
	}
}
