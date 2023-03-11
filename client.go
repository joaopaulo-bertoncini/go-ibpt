package ibpt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultEndpoint contains endpoint URL of FCM service.
	DefaultEndpointProducts = "https://apidoni.ibpt.org.br/api/v1/produtos"

	DefaultEndpointServices = "https://apidoni.ibpt.org.br/api/v1/servicos"

	// DefaultTimeout duration in second
	DefaultTimeout time.Duration = 30 * time.Second

	Service string = "S"
	Product string = "P"
)

// Client abstracts the interaction between the application server and the
// IBPT server via HTTP protocol.
// If the `HTTP` field is nil, a zeroed http.Client will be allocated and used
// to send messages.
type Client struct {
	client   *http.Client
	endpoint string
	timeout  time.Duration
	kind     string
}

// NewClient creates new IBPT Client based on API key and
// with default endpoint Products and http client.
func NewClientProduct(opts ...Option) (*Client, error) {
	c := &Client{
		endpoint: DefaultEndpointProducts,
		client:   &http.Client{},
		timeout:  DefaultTimeout,
		kind:     Product,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// NewClient creates new IBPT Client based on API key and
// with default endpoint Services and http client.
func NewClientService(opts ...Option) (*Client, error) {
	c := &Client{
		endpoint: DefaultEndpointServices,
		client:   &http.Client{},
		timeout:  DefaultTimeout,
		kind:     Service,
	}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// SendWithContext sends a message to the IBPT server without retrying in case of service
// unavailability. A non-nil error is returned if a non-recoverable error
// occurs (i.e. if the response status is not "200 OK").
// Behaves just like regular send, but uses external context.
func (c *Client) SendWithContext(ctx context.Context, data *Request) (*Response, error) {
	// validate
	if err := data.Validate(); err != nil {
		return nil, err
	}

	return c.send(ctx, data)
}

// Send sends a Request to the IBPT server without retrying in case of service
// unavailability. A non-nil error is returned if a non-recoverable error
// occurs (i.e. if the response status is not "200 OK").
func (c *Client) Send(data *Request) (*Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.SendWithContext(ctx, data)
}

// SendWithRetry sends a Request to the IBPT server with defined number of
// retrying in case of temporary error.
func (c *Client) SendWithRetry(data *Request, retryAttempts int) (*Response, error) {
	return c.SendWithRetryWithContext(context.Background(), data, retryAttempts)
}

// SendWithRetryWithContext sends a Request to the IBPT server with defined number of
// retrying in case of temporary error.
// Behaves just like regular SendWithRetry, but uses external context.
func (c *Client) SendWithRetryWithContext(ctx context.Context, data *Request, retryAttempts int) (*Response, error) {
	// validate
	if err := data.Validate(); err != nil {
		return nil, err
	}

	resp := new(Response)
	err := retry(func() error {
		ctx, cancel := context.WithTimeout(ctx, c.timeout)
		defer cancel()
		var er error
		resp, er = c.send(ctx, data)
		return er
	}, retryAttempts)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// send sends a request.
func (c *Client) send(ctx context.Context, data *Request) (*Response, error) {
	// create request
	req, err := http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	// add headers
	req.Header.Add("Content-Type", "application/json")

	// appending to existing query args
	params := req.URL.Query()
	params.Add("token", data.Token)
	params.Add("cnpj", data.CNPJ)
	params.Add("codigo", data.Code)
	params.Add("uf", data.UF)
	params.Add("ex", fmt.Sprintf("%d", data.EX))
	params.Add("codigoInterno", data.InternalCode)
	params.Add("unidadeMedida", data.UnitMeasurement)
	params.Add("descricao", data.Description)
	params.Add("valor", fmt.Sprintf("%2.f", data.Value))
	params.Add("gtin", data.Gtin)

	// assign encoded query string to http request
	req.URL.RawQuery = params.Encode()
	// execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.New("Error when sending request to the server")
	}
	defer resp.Body.Close()

	// check response status
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, errors.New(fmt.Sprintf("%d error: %s", resp.StatusCode, resp.Status))
		}
		return nil, fmt.Errorf("%d error: %s", resp.StatusCode, resp.Status)
	}

	response := new(Response)
	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return nil, err
	}

	return response, nil
}
