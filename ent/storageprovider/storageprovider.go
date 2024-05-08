// Code generated by ent, DO NOT EDIT.

package storageprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the storageprovider type in the database.
	Label = "storage_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBucket holds the string denoting the bucket field in the database.
	FieldBucket = "bucket"
	// FieldSecretID holds the string denoting the secret_id field in the database.
	FieldSecretID = "secret_id"
	// FieldSecretKey holds the string denoting the secret_key field in the database.
	FieldSecretKey = "secret_key"
	// FieldEndpoint holds the string denoting the endpoint field in the database.
	FieldEndpoint = "endpoint"
	// FieldPreviewHost holds the string denoting the preview_host field in the database.
	FieldPreviewHost = "preview_host"
	// FieldFolder holds the string denoting the folder field in the database.
	FieldFolder = "folder"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldIsDefault holds the string denoting the is_default field in the database.
	FieldIsDefault = "is_default"
	// EdgeCloudfiles holds the string denoting the cloudfiles edge name in mutations.
	EdgeCloudfiles = "cloudfiles"
	// Table holds the table name of the storageprovider in the database.
	Table = "fms_storage_providers"
	// CloudfilesTable is the table that holds the cloudfiles relation/edge.
	CloudfilesTable = "fms_cloud_files"
	// CloudfilesInverseTable is the table name for the CloudFile entity.
	// It exists in this package in order to avoid circular dependency with the "cloudfile" package.
	CloudfilesInverseTable = "fms_cloud_files"
	// CloudfilesColumn is the table column denoting the cloudfiles relation/edge.
	CloudfilesColumn = "cloud_file_storage_providers"
)

// Columns holds all SQL columns for storageprovider fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldState,
	FieldName,
	FieldBucket,
	FieldSecretID,
	FieldSecretKey,
	FieldEndpoint,
	FieldPreviewHost,
	FieldFolder,
	FieldRegion,
	FieldIsDefault,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState bool
	// DefaultIsDefault holds the default value on creation for the "is_default" field.
	DefaultIsDefault bool
)

// OrderOption defines the ordering options for the StorageProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByBucket orders the results by the bucket field.
func ByBucket(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBucket, opts...).ToFunc()
}

// BySecretID orders the results by the secret_id field.
func BySecretID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretID, opts...).ToFunc()
}

// BySecretKey orders the results by the secret_key field.
func BySecretKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSecretKey, opts...).ToFunc()
}

// ByEndpoint orders the results by the endpoint field.
func ByEndpoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndpoint, opts...).ToFunc()
}

// ByPreviewHost orders the results by the preview_host field.
func ByPreviewHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPreviewHost, opts...).ToFunc()
}

// ByFolder orders the results by the folder field.
func ByFolder(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFolder, opts...).ToFunc()
}

// ByRegion orders the results by the region field.
func ByRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRegion, opts...).ToFunc()
}

// ByIsDefault orders the results by the is_default field.
func ByIsDefault(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsDefault, opts...).ToFunc()
}

// ByCloudfilesCount orders the results by cloudfiles count.
func ByCloudfilesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCloudfilesStep(), opts...)
	}
}

// ByCloudfiles orders the results by cloudfiles terms.
func ByCloudfiles(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCloudfilesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCloudfilesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CloudfilesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CloudfilesTable, CloudfilesColumn),
	)
}
