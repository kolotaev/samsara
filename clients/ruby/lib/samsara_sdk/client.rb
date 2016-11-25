require 'samsara_sdk/config'
require 'samsara_sdk/publisher'
require 'samsara_sdk/ring_buffer'
require 'samsara_sdk/event'

# Samsara SDK
module SamsaraSDK
  # A client for ingesting events into Samsara.
  # It is the main interface to communicate with Samsara API.
  class Client
    # initialize
    #
    # @param config [Hash] Configuration overrides.
    # @raise [SamsaraSDK::ConfigValidationError] if any config option is incorrect.
    def initialize(config)
      Config.setup! config
      @publisher = Publisher.new
      @queue = RingBuffer.new(Config.get[:max_buffer_size])
    end

    # Publishes given events list to Ingestion API immediately.
    #
    # @param events [Array<Hash>] List of events.
    # @return [Boolean] The result of publish operation.
    # @raise [SamsaraSDK::EventValidationError] if any option of the given event is invalid.
    # @see http://samsara-analytics.io/docs/design/events-spec Event specification
    def publish_events(events)
      events = events.map { |event| Event.validate(Event.enrich(event)) }
      @publisher.post events
    end

    # Pushes event to internal events' queue.
    #
    # @param event [Hash] Event data.
    # @option data [String] :eventName Name of the event.
    # @option data [String] :sourceId Source ID of the event.
    # @option data [Integer] :timestamp Timestamp in milliseconds.
    # @see http://samsara-analytics.io/docs/design/events-spec Event specification
    def record_event(event)
      # Event.enrich event
      # Event.validate event
      # @queue << event
    end
  end
end
