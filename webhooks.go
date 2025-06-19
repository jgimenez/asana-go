package asana

// WebhookBase objects represent the state of an active subscription for a server to be updated with information from Asana.
// This schema represents the subscription itself, not the objects that are sent to the server.
// For information on those please refer to the Event schema.
type Webhook struct {
	WebhookBase
	// Read-only. Globally unique ID of the object
	ID       string `json:"gid,omitempty"`
	Resource struct {
		ID           string `json:"gid,omitempty"`
		Name         string `json:"name,omitempty"`
		ResourceType string `json:"resource_type,omitempty"`
	} `json:"resource,omitempty"`
}

type WebhookBase struct {
	// Whitelist of filters to apply to events from this webhook.
	Filters []WebhookFilters `json:"filters,omitempty"`

	// A generic Asana Resource, containing a globally unique identifier.
	Resource string `json:"resource,omitempty"`

	// The URL to receive the HTTP POST.
	Target string `json:"target,omitempty"`
}

type WebhookFilters struct {
	// The type of change on the resource to pass through the filter.
	Action string `json:"action,omitempty"`
	// Conditional. A whitelist of fields for events which will pass the filter when the resource is changed.
	Fields string `json:"fields,omitempty"`
	// The resource subtype of the resource that the filter applies to.
	ResourceSubtype string `json:"resource_subtype,omitempty"`
	// The resource type of the resource that the filter applies to.
	ResourceType string `json:"resource_type,omitempty"`
}

// CreateWebhook Establish a webhook
func (c *Client) CreateWebhook(webhook *WebhookBase, options ...*Options) (*Webhook, error) {
	c.trace("Creating webhook...\n")
	result := &Webhook{}

	err := c.post("/webhooks", webhook, result, options...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateWebhook Establish a webhook
func (c *Client) GetWebhooks(options ...*Options) (*[]Webhook, *NextPage, error) {
	c.trace("Listing webhooks...\n")
	result := make([]Webhook, 0)

	nextPage, err := c.get("/webhooks", nil, &result, options...)
	if err != nil {
		return nil, nil, err
	}

	return &result, nextPage, nil
}
