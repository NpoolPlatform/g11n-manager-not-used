// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/g11n-manager/pkg/db/ent/country"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Country is the client for interacting with the Country builders.
	Country *CountryClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Country = NewCountryClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Country: NewCountryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Country: NewCountryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Country.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Country.Use(hooks...)
}

// CountryClient is a client for the Country schema.
type CountryClient struct {
	config
}

// NewCountryClient returns a client for the Country from the given config.
func NewCountryClient(c config) *CountryClient {
	return &CountryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `country.Hooks(f(g(h())))`.
func (c *CountryClient) Use(hooks ...Hook) {
	c.hooks.Country = append(c.hooks.Country, hooks...)
}

// Create returns a builder for creating a Country entity.
func (c *CountryClient) Create() *CountryCreate {
	mutation := newCountryMutation(c.config, OpCreate)
	return &CountryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Country entities.
func (c *CountryClient) CreateBulk(builders ...*CountryCreate) *CountryCreateBulk {
	return &CountryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Country.
func (c *CountryClient) Update() *CountryUpdate {
	mutation := newCountryMutation(c.config, OpUpdate)
	return &CountryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CountryClient) UpdateOne(co *Country) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountry(co))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CountryClient) UpdateOneID(id uuid.UUID) *CountryUpdateOne {
	mutation := newCountryMutation(c.config, OpUpdateOne, withCountryID(id))
	return &CountryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Country.
func (c *CountryClient) Delete() *CountryDelete {
	mutation := newCountryMutation(c.config, OpDelete)
	return &CountryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CountryClient) DeleteOne(co *Country) *CountryDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CountryClient) DeleteOneID(id uuid.UUID) *CountryDeleteOne {
	builder := c.Delete().Where(country.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CountryDeleteOne{builder}
}

// Query returns a query builder for Country.
func (c *CountryClient) Query() *CountryQuery {
	return &CountryQuery{
		config: c.config,
	}
}

// Get returns a Country entity by its id.
func (c *CountryClient) Get(ctx context.Context, id uuid.UUID) (*Country, error) {
	return c.Query().Where(country.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CountryClient) GetX(ctx context.Context, id uuid.UUID) *Country {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CountryClient) Hooks() []Hook {
	hooks := c.hooks.Country
	return append(hooks[:len(hooks):len(hooks)], country.Hooks[:]...)
}
