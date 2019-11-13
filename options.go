package tipe

type Option func(*APIClient)

// Set the host on the client
func Host(host string) Option {
	return func(c *APIClient) {
		c.host = host
	}
}

// Set the key on the client
func Key(key string) Option {
	return func(c *APIClient) {
		c.key = key
	}
}

// Set offline mode
func Offline(offline bool) Option {
	return func(c *APIClient) {
		c.offline = offline
	}
}

// Set the port in offline mode
func Port(port int) Option {
	return func(c *APIClient) {
		c.port = port
	}
}

// Set the project
func Project(project string) Option {
	return func(c *APIClient) {
		c.project = project
	}
}
