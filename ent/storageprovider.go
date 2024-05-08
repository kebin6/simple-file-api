// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kebin6/simple-file-api/ent/storageprovider"
)

// StorageProvider is the model entity for the StorageProvider schema.
type StorageProvider struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// State true: normal false: ban | 状态 true 正常 false 禁用
	State bool `json:"state,omitempty"`
	// The cloud storage service name | 服务名称
	Name string `json:"name,omitempty"`
	// The cloud storage bucket name | 云存储服务的存储桶
	Bucket string `json:"bucket,omitempty"`
	// The secret ID | 密钥 ID
	SecretID string `json:"secret_id,omitempty"`
	// The secret key | 密钥 Key
	SecretKey string `json:"secret_key,omitempty"`
	// The service URL | 服务器地址
	Endpoint string `json:"endpoint,omitempty"`
	// The file preview Host | 文件的预览地址
	PreviewHost string `json:"preview_host,omitempty"`
	// The folder in cloud | 云服务目标文件夹
	Folder string `json:"folder,omitempty"`
	// The service region | 服务器所在地区
	Region string `json:"region,omitempty"`
	// Is it the default provider | 是否为默认提供商
	IsDefault bool `json:"is_default,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StorageProviderQuery when eager-loading is set.
	Edges        StorageProviderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// StorageProviderEdges holds the relations/edges for other nodes in the graph.
type StorageProviderEdges struct {
	// Cloudfiles holds the value of the cloudfiles edge.
	Cloudfiles []*CloudFile `json:"cloudfiles,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CloudfilesOrErr returns the Cloudfiles value or an error if the edge
// was not loaded in eager-loading.
func (e StorageProviderEdges) CloudfilesOrErr() ([]*CloudFile, error) {
	if e.loadedTypes[0] {
		return e.Cloudfiles, nil
	}
	return nil, &NotLoadedError{edge: "cloudfiles"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*StorageProvider) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case storageprovider.FieldState, storageprovider.FieldIsDefault:
			values[i] = new(sql.NullBool)
		case storageprovider.FieldID:
			values[i] = new(sql.NullInt64)
		case storageprovider.FieldName, storageprovider.FieldBucket, storageprovider.FieldSecretID, storageprovider.FieldSecretKey, storageprovider.FieldEndpoint, storageprovider.FieldPreviewHost, storageprovider.FieldFolder, storageprovider.FieldRegion:
			values[i] = new(sql.NullString)
		case storageprovider.FieldCreatedAt, storageprovider.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the StorageProvider fields.
func (sp *StorageProvider) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case storageprovider.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sp.ID = uint64(value.Int64)
		case storageprovider.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sp.CreatedAt = value.Time
			}
		case storageprovider.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sp.UpdatedAt = value.Time
			}
		case storageprovider.FieldState:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				sp.State = value.Bool
			}
		case storageprovider.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				sp.Name = value.String
			}
		case storageprovider.FieldBucket:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field bucket", values[i])
			} else if value.Valid {
				sp.Bucket = value.String
			}
		case storageprovider.FieldSecretID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_id", values[i])
			} else if value.Valid {
				sp.SecretID = value.String
			}
		case storageprovider.FieldSecretKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret_key", values[i])
			} else if value.Valid {
				sp.SecretKey = value.String
			}
		case storageprovider.FieldEndpoint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field endpoint", values[i])
			} else if value.Valid {
				sp.Endpoint = value.String
			}
		case storageprovider.FieldPreviewHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preview_host", values[i])
			} else if value.Valid {
				sp.PreviewHost = value.String
			}
		case storageprovider.FieldFolder:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field folder", values[i])
			} else if value.Valid {
				sp.Folder = value.String
			}
		case storageprovider.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				sp.Region = value.String
			}
		case storageprovider.FieldIsDefault:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_default", values[i])
			} else if value.Valid {
				sp.IsDefault = value.Bool
			}
		default:
			sp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the StorageProvider.
// This includes values selected through modifiers, order, etc.
func (sp *StorageProvider) Value(name string) (ent.Value, error) {
	return sp.selectValues.Get(name)
}

// QueryCloudfiles queries the "cloudfiles" edge of the StorageProvider entity.
func (sp *StorageProvider) QueryCloudfiles() *CloudFileQuery {
	return NewStorageProviderClient(sp.config).QueryCloudfiles(sp)
}

// Update returns a builder for updating this StorageProvider.
// Note that you need to call StorageProvider.Unwrap() before calling this method if this StorageProvider
// was returned from a transaction, and the transaction was committed or rolled back.
func (sp *StorageProvider) Update() *StorageProviderUpdateOne {
	return NewStorageProviderClient(sp.config).UpdateOne(sp)
}

// Unwrap unwraps the StorageProvider entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sp *StorageProvider) Unwrap() *StorageProvider {
	_tx, ok := sp.config.driver.(*txDriver)
	if !ok {
		panic("ent: StorageProvider is not a transactional entity")
	}
	sp.config.driver = _tx.drv
	return sp
}

// String implements the fmt.Stringer.
func (sp *StorageProvider) String() string {
	var builder strings.Builder
	builder.WriteString("StorageProvider(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", sp.State))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(sp.Name)
	builder.WriteString(", ")
	builder.WriteString("bucket=")
	builder.WriteString(sp.Bucket)
	builder.WriteString(", ")
	builder.WriteString("secret_id=")
	builder.WriteString(sp.SecretID)
	builder.WriteString(", ")
	builder.WriteString("secret_key=")
	builder.WriteString(sp.SecretKey)
	builder.WriteString(", ")
	builder.WriteString("endpoint=")
	builder.WriteString(sp.Endpoint)
	builder.WriteString(", ")
	builder.WriteString("preview_host=")
	builder.WriteString(sp.PreviewHost)
	builder.WriteString(", ")
	builder.WriteString("folder=")
	builder.WriteString(sp.Folder)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(sp.Region)
	builder.WriteString(", ")
	builder.WriteString("is_default=")
	builder.WriteString(fmt.Sprintf("%v", sp.IsDefault))
	builder.WriteByte(')')
	return builder.String()
}

// StorageProviders is a parsable slice of StorageProvider.
type StorageProviders []*StorageProvider
