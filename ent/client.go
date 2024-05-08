// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	uuid "github.com/gofrs/uuid/v5"
	"github.com/kebin6/simple-file-api/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kebin6/simple-file-api/ent/cloudfile"
	"github.com/kebin6/simple-file-api/ent/cloudfiletag"
	"github.com/kebin6/simple-file-api/ent/file"
	"github.com/kebin6/simple-file-api/ent/filetag"
	"github.com/kebin6/simple-file-api/ent/storageprovider"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CloudFile is the client for interacting with the CloudFile builders.
	CloudFile *CloudFileClient
	// CloudFileTag is the client for interacting with the CloudFileTag builders.
	CloudFileTag *CloudFileTagClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// FileTag is the client for interacting with the FileTag builders.
	FileTag *FileTagClient
	// StorageProvider is the client for interacting with the StorageProvider builders.
	StorageProvider *StorageProviderClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CloudFile = NewCloudFileClient(c.config)
	c.CloudFileTag = NewCloudFileTagClient(c.config)
	c.File = NewFileClient(c.config)
	c.FileTag = NewFileTagClient(c.config)
	c.StorageProvider = NewStorageProviderClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:             ctx,
		config:          cfg,
		CloudFile:       NewCloudFileClient(cfg),
		CloudFileTag:    NewCloudFileTagClient(cfg),
		File:            NewFileClient(cfg),
		FileTag:         NewFileTagClient(cfg),
		StorageProvider: NewStorageProviderClient(cfg),
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
		ctx:             ctx,
		config:          cfg,
		CloudFile:       NewCloudFileClient(cfg),
		CloudFileTag:    NewCloudFileTagClient(cfg),
		File:            NewFileClient(cfg),
		FileTag:         NewFileTagClient(cfg),
		StorageProvider: NewStorageProviderClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CloudFile.
//		Query().
//		Count(ctx)
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
	c.CloudFile.Use(hooks...)
	c.CloudFileTag.Use(hooks...)
	c.File.Use(hooks...)
	c.FileTag.Use(hooks...)
	c.StorageProvider.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.CloudFile.Intercept(interceptors...)
	c.CloudFileTag.Intercept(interceptors...)
	c.File.Intercept(interceptors...)
	c.FileTag.Intercept(interceptors...)
	c.StorageProvider.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CloudFileMutation:
		return c.CloudFile.mutate(ctx, m)
	case *CloudFileTagMutation:
		return c.CloudFileTag.mutate(ctx, m)
	case *FileMutation:
		return c.File.mutate(ctx, m)
	case *FileTagMutation:
		return c.FileTag.mutate(ctx, m)
	case *StorageProviderMutation:
		return c.StorageProvider.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CloudFileClient is a client for the CloudFile schema.
type CloudFileClient struct {
	config
}

// NewCloudFileClient returns a client for the CloudFile from the given config.
func NewCloudFileClient(c config) *CloudFileClient {
	return &CloudFileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cloudfile.Hooks(f(g(h())))`.
func (c *CloudFileClient) Use(hooks ...Hook) {
	c.hooks.CloudFile = append(c.hooks.CloudFile, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cloudfile.Intercept(f(g(h())))`.
func (c *CloudFileClient) Intercept(interceptors ...Interceptor) {
	c.inters.CloudFile = append(c.inters.CloudFile, interceptors...)
}

// Create returns a builder for creating a CloudFile entity.
func (c *CloudFileClient) Create() *CloudFileCreate {
	mutation := newCloudFileMutation(c.config, OpCreate)
	return &CloudFileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CloudFile entities.
func (c *CloudFileClient) CreateBulk(builders ...*CloudFileCreate) *CloudFileCreateBulk {
	return &CloudFileCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CloudFileClient) MapCreateBulk(slice any, setFunc func(*CloudFileCreate, int)) *CloudFileCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CloudFileCreateBulk{err: fmt.Errorf("calling to CloudFileClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CloudFileCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CloudFileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CloudFile.
func (c *CloudFileClient) Update() *CloudFileUpdate {
	mutation := newCloudFileMutation(c.config, OpUpdate)
	return &CloudFileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CloudFileClient) UpdateOne(cf *CloudFile) *CloudFileUpdateOne {
	mutation := newCloudFileMutation(c.config, OpUpdateOne, withCloudFile(cf))
	return &CloudFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CloudFileClient) UpdateOneID(id uuid.UUID) *CloudFileUpdateOne {
	mutation := newCloudFileMutation(c.config, OpUpdateOne, withCloudFileID(id))
	return &CloudFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CloudFile.
func (c *CloudFileClient) Delete() *CloudFileDelete {
	mutation := newCloudFileMutation(c.config, OpDelete)
	return &CloudFileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CloudFileClient) DeleteOne(cf *CloudFile) *CloudFileDeleteOne {
	return c.DeleteOneID(cf.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CloudFileClient) DeleteOneID(id uuid.UUID) *CloudFileDeleteOne {
	builder := c.Delete().Where(cloudfile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CloudFileDeleteOne{builder}
}

// Query returns a query builder for CloudFile.
func (c *CloudFileClient) Query() *CloudFileQuery {
	return &CloudFileQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCloudFile},
		inters: c.Interceptors(),
	}
}

// Get returns a CloudFile entity by its id.
func (c *CloudFileClient) Get(ctx context.Context, id uuid.UUID) (*CloudFile, error) {
	return c.Query().Where(cloudfile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CloudFileClient) GetX(ctx context.Context, id uuid.UUID) *CloudFile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStorageProviders queries the storage_providers edge of a CloudFile.
func (c *CloudFileClient) QueryStorageProviders(cf *CloudFile) *StorageProviderQuery {
	query := (&StorageProviderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cf.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cloudfile.Table, cloudfile.FieldID, id),
			sqlgraph.To(storageprovider.Table, storageprovider.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, cloudfile.StorageProvidersTable, cloudfile.StorageProvidersColumn),
		)
		fromV = sqlgraph.Neighbors(cf.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTags queries the tags edge of a CloudFile.
func (c *CloudFileClient) QueryTags(cf *CloudFile) *CloudFileTagQuery {
	query := (&CloudFileTagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cf.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cloudfile.Table, cloudfile.FieldID, id),
			sqlgraph.To(cloudfiletag.Table, cloudfiletag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, cloudfile.TagsTable, cloudfile.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cf.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CloudFileClient) Hooks() []Hook {
	return c.hooks.CloudFile
}

// Interceptors returns the client interceptors.
func (c *CloudFileClient) Interceptors() []Interceptor {
	return c.inters.CloudFile
}

func (c *CloudFileClient) mutate(ctx context.Context, m *CloudFileMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CloudFileCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CloudFileUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CloudFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CloudFileDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CloudFile mutation op: %q", m.Op())
	}
}

// CloudFileTagClient is a client for the CloudFileTag schema.
type CloudFileTagClient struct {
	config
}

// NewCloudFileTagClient returns a client for the CloudFileTag from the given config.
func NewCloudFileTagClient(c config) *CloudFileTagClient {
	return &CloudFileTagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cloudfiletag.Hooks(f(g(h())))`.
func (c *CloudFileTagClient) Use(hooks ...Hook) {
	c.hooks.CloudFileTag = append(c.hooks.CloudFileTag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cloudfiletag.Intercept(f(g(h())))`.
func (c *CloudFileTagClient) Intercept(interceptors ...Interceptor) {
	c.inters.CloudFileTag = append(c.inters.CloudFileTag, interceptors...)
}

// Create returns a builder for creating a CloudFileTag entity.
func (c *CloudFileTagClient) Create() *CloudFileTagCreate {
	mutation := newCloudFileTagMutation(c.config, OpCreate)
	return &CloudFileTagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CloudFileTag entities.
func (c *CloudFileTagClient) CreateBulk(builders ...*CloudFileTagCreate) *CloudFileTagCreateBulk {
	return &CloudFileTagCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CloudFileTagClient) MapCreateBulk(slice any, setFunc func(*CloudFileTagCreate, int)) *CloudFileTagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CloudFileTagCreateBulk{err: fmt.Errorf("calling to CloudFileTagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CloudFileTagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CloudFileTagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CloudFileTag.
func (c *CloudFileTagClient) Update() *CloudFileTagUpdate {
	mutation := newCloudFileTagMutation(c.config, OpUpdate)
	return &CloudFileTagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CloudFileTagClient) UpdateOne(cft *CloudFileTag) *CloudFileTagUpdateOne {
	mutation := newCloudFileTagMutation(c.config, OpUpdateOne, withCloudFileTag(cft))
	return &CloudFileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CloudFileTagClient) UpdateOneID(id uint64) *CloudFileTagUpdateOne {
	mutation := newCloudFileTagMutation(c.config, OpUpdateOne, withCloudFileTagID(id))
	return &CloudFileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CloudFileTag.
func (c *CloudFileTagClient) Delete() *CloudFileTagDelete {
	mutation := newCloudFileTagMutation(c.config, OpDelete)
	return &CloudFileTagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CloudFileTagClient) DeleteOne(cft *CloudFileTag) *CloudFileTagDeleteOne {
	return c.DeleteOneID(cft.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CloudFileTagClient) DeleteOneID(id uint64) *CloudFileTagDeleteOne {
	builder := c.Delete().Where(cloudfiletag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CloudFileTagDeleteOne{builder}
}

// Query returns a query builder for CloudFileTag.
func (c *CloudFileTagClient) Query() *CloudFileTagQuery {
	return &CloudFileTagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCloudFileTag},
		inters: c.Interceptors(),
	}
}

// Get returns a CloudFileTag entity by its id.
func (c *CloudFileTagClient) Get(ctx context.Context, id uint64) (*CloudFileTag, error) {
	return c.Query().Where(cloudfiletag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CloudFileTagClient) GetX(ctx context.Context, id uint64) *CloudFileTag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCloudFiles queries the cloud_files edge of a CloudFileTag.
func (c *CloudFileTagClient) QueryCloudFiles(cft *CloudFileTag) *CloudFileQuery {
	query := (&CloudFileClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cft.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cloudfiletag.Table, cloudfiletag.FieldID, id),
			sqlgraph.To(cloudfile.Table, cloudfile.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, cloudfiletag.CloudFilesTable, cloudfiletag.CloudFilesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cft.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CloudFileTagClient) Hooks() []Hook {
	return c.hooks.CloudFileTag
}

// Interceptors returns the client interceptors.
func (c *CloudFileTagClient) Interceptors() []Interceptor {
	return c.inters.CloudFileTag
}

func (c *CloudFileTagClient) mutate(ctx context.Context, m *CloudFileTagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CloudFileTagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CloudFileTagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CloudFileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CloudFileTagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CloudFileTag mutation op: %q", m.Op())
	}
}

// FileClient is a client for the File schema.
type FileClient struct {
	config
}

// NewFileClient returns a client for the File from the given config.
func NewFileClient(c config) *FileClient {
	return &FileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `file.Hooks(f(g(h())))`.
func (c *FileClient) Use(hooks ...Hook) {
	c.hooks.File = append(c.hooks.File, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `file.Intercept(f(g(h())))`.
func (c *FileClient) Intercept(interceptors ...Interceptor) {
	c.inters.File = append(c.inters.File, interceptors...)
}

// Create returns a builder for creating a File entity.
func (c *FileClient) Create() *FileCreate {
	mutation := newFileMutation(c.config, OpCreate)
	return &FileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of File entities.
func (c *FileClient) CreateBulk(builders ...*FileCreate) *FileCreateBulk {
	return &FileCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FileClient) MapCreateBulk(slice any, setFunc func(*FileCreate, int)) *FileCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FileCreateBulk{err: fmt.Errorf("calling to FileClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FileCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for File.
func (c *FileClient) Update() *FileUpdate {
	mutation := newFileMutation(c.config, OpUpdate)
	return &FileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileClient) UpdateOne(f *File) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFile(f))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileClient) UpdateOneID(id uuid.UUID) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFileID(id))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for File.
func (c *FileClient) Delete() *FileDelete {
	mutation := newFileMutation(c.config, OpDelete)
	return &FileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FileClient) DeleteOne(f *File) *FileDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FileClient) DeleteOneID(id uuid.UUID) *FileDeleteOne {
	builder := c.Delete().Where(file.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileDeleteOne{builder}
}

// Query returns a query builder for File.
func (c *FileClient) Query() *FileQuery {
	return &FileQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFile},
		inters: c.Interceptors(),
	}
}

// Get returns a File entity by its id.
func (c *FileClient) Get(ctx context.Context, id uuid.UUID) (*File, error) {
	return c.Query().Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileClient) GetX(ctx context.Context, id uuid.UUID) *File {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a File.
func (c *FileClient) QueryTags(f *File) *FileTagQuery {
	query := (&FileTagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, id),
			sqlgraph.To(filetag.Table, filetag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, file.TagsTable, file.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FileClient) Hooks() []Hook {
	return c.hooks.File
}

// Interceptors returns the client interceptors.
func (c *FileClient) Interceptors() []Interceptor {
	return c.inters.File
}

func (c *FileClient) mutate(ctx context.Context, m *FileMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FileCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FileUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FileDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown File mutation op: %q", m.Op())
	}
}

// FileTagClient is a client for the FileTag schema.
type FileTagClient struct {
	config
}

// NewFileTagClient returns a client for the FileTag from the given config.
func NewFileTagClient(c config) *FileTagClient {
	return &FileTagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `filetag.Hooks(f(g(h())))`.
func (c *FileTagClient) Use(hooks ...Hook) {
	c.hooks.FileTag = append(c.hooks.FileTag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `filetag.Intercept(f(g(h())))`.
func (c *FileTagClient) Intercept(interceptors ...Interceptor) {
	c.inters.FileTag = append(c.inters.FileTag, interceptors...)
}

// Create returns a builder for creating a FileTag entity.
func (c *FileTagClient) Create() *FileTagCreate {
	mutation := newFileTagMutation(c.config, OpCreate)
	return &FileTagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FileTag entities.
func (c *FileTagClient) CreateBulk(builders ...*FileTagCreate) *FileTagCreateBulk {
	return &FileTagCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FileTagClient) MapCreateBulk(slice any, setFunc func(*FileTagCreate, int)) *FileTagCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FileTagCreateBulk{err: fmt.Errorf("calling to FileTagClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FileTagCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FileTagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FileTag.
func (c *FileTagClient) Update() *FileTagUpdate {
	mutation := newFileTagMutation(c.config, OpUpdate)
	return &FileTagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileTagClient) UpdateOne(ft *FileTag) *FileTagUpdateOne {
	mutation := newFileTagMutation(c.config, OpUpdateOne, withFileTag(ft))
	return &FileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileTagClient) UpdateOneID(id uint64) *FileTagUpdateOne {
	mutation := newFileTagMutation(c.config, OpUpdateOne, withFileTagID(id))
	return &FileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FileTag.
func (c *FileTagClient) Delete() *FileTagDelete {
	mutation := newFileTagMutation(c.config, OpDelete)
	return &FileTagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FileTagClient) DeleteOne(ft *FileTag) *FileTagDeleteOne {
	return c.DeleteOneID(ft.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FileTagClient) DeleteOneID(id uint64) *FileTagDeleteOne {
	builder := c.Delete().Where(filetag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileTagDeleteOne{builder}
}

// Query returns a query builder for FileTag.
func (c *FileTagClient) Query() *FileTagQuery {
	return &FileTagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFileTag},
		inters: c.Interceptors(),
	}
}

// Get returns a FileTag entity by its id.
func (c *FileTagClient) Get(ctx context.Context, id uint64) (*FileTag, error) {
	return c.Query().Where(filetag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileTagClient) GetX(ctx context.Context, id uint64) *FileTag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFiles queries the files edge of a FileTag.
func (c *FileTagClient) QueryFiles(ft *FileTag) *FileQuery {
	query := (&FileClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ft.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(filetag.Table, filetag.FieldID, id),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, filetag.FilesTable, filetag.FilesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(ft.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FileTagClient) Hooks() []Hook {
	return c.hooks.FileTag
}

// Interceptors returns the client interceptors.
func (c *FileTagClient) Interceptors() []Interceptor {
	return c.inters.FileTag
}

func (c *FileTagClient) mutate(ctx context.Context, m *FileTagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FileTagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FileTagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FileTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FileTagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown FileTag mutation op: %q", m.Op())
	}
}

// StorageProviderClient is a client for the StorageProvider schema.
type StorageProviderClient struct {
	config
}

// NewStorageProviderClient returns a client for the StorageProvider from the given config.
func NewStorageProviderClient(c config) *StorageProviderClient {
	return &StorageProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `storageprovider.Hooks(f(g(h())))`.
func (c *StorageProviderClient) Use(hooks ...Hook) {
	c.hooks.StorageProvider = append(c.hooks.StorageProvider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `storageprovider.Intercept(f(g(h())))`.
func (c *StorageProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.StorageProvider = append(c.inters.StorageProvider, interceptors...)
}

// Create returns a builder for creating a StorageProvider entity.
func (c *StorageProviderClient) Create() *StorageProviderCreate {
	mutation := newStorageProviderMutation(c.config, OpCreate)
	return &StorageProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StorageProvider entities.
func (c *StorageProviderClient) CreateBulk(builders ...*StorageProviderCreate) *StorageProviderCreateBulk {
	return &StorageProviderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StorageProviderClient) MapCreateBulk(slice any, setFunc func(*StorageProviderCreate, int)) *StorageProviderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StorageProviderCreateBulk{err: fmt.Errorf("calling to StorageProviderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StorageProviderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StorageProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StorageProvider.
func (c *StorageProviderClient) Update() *StorageProviderUpdate {
	mutation := newStorageProviderMutation(c.config, OpUpdate)
	return &StorageProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StorageProviderClient) UpdateOne(sp *StorageProvider) *StorageProviderUpdateOne {
	mutation := newStorageProviderMutation(c.config, OpUpdateOne, withStorageProvider(sp))
	return &StorageProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StorageProviderClient) UpdateOneID(id uint64) *StorageProviderUpdateOne {
	mutation := newStorageProviderMutation(c.config, OpUpdateOne, withStorageProviderID(id))
	return &StorageProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StorageProvider.
func (c *StorageProviderClient) Delete() *StorageProviderDelete {
	mutation := newStorageProviderMutation(c.config, OpDelete)
	return &StorageProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StorageProviderClient) DeleteOne(sp *StorageProvider) *StorageProviderDeleteOne {
	return c.DeleteOneID(sp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StorageProviderClient) DeleteOneID(id uint64) *StorageProviderDeleteOne {
	builder := c.Delete().Where(storageprovider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StorageProviderDeleteOne{builder}
}

// Query returns a query builder for StorageProvider.
func (c *StorageProviderClient) Query() *StorageProviderQuery {
	return &StorageProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStorageProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a StorageProvider entity by its id.
func (c *StorageProviderClient) Get(ctx context.Context, id uint64) (*StorageProvider, error) {
	return c.Query().Where(storageprovider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StorageProviderClient) GetX(ctx context.Context, id uint64) *StorageProvider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCloudfiles queries the cloudfiles edge of a StorageProvider.
func (c *StorageProviderClient) QueryCloudfiles(sp *StorageProvider) *CloudFileQuery {
	query := (&CloudFileClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := sp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(storageprovider.Table, storageprovider.FieldID, id),
			sqlgraph.To(cloudfile.Table, cloudfile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, storageprovider.CloudfilesTable, storageprovider.CloudfilesColumn),
		)
		fromV = sqlgraph.Neighbors(sp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StorageProviderClient) Hooks() []Hook {
	return c.hooks.StorageProvider
}

// Interceptors returns the client interceptors.
func (c *StorageProviderClient) Interceptors() []Interceptor {
	return c.inters.StorageProvider
}

func (c *StorageProviderClient) mutate(ctx context.Context, m *StorageProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StorageProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StorageProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StorageProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StorageProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown StorageProvider mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		CloudFile, CloudFileTag, File, FileTag, StorageProvider []ent.Hook
	}
	inters struct {
		CloudFile, CloudFileTag, File, FileTag, StorageProvider []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}