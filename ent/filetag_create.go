// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/kebin6/simple-file-api/ent/file"
	"github.com/kebin6/simple-file-api/ent/filetag"
)

// FileTagCreate is the builder for creating a FileTag entity.
type FileTagCreate struct {
	config
	mutation *FileTagMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ftc *FileTagCreate) SetCreatedAt(t time.Time) *FileTagCreate {
	ftc.mutation.SetCreatedAt(t)
	return ftc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ftc *FileTagCreate) SetNillableCreatedAt(t *time.Time) *FileTagCreate {
	if t != nil {
		ftc.SetCreatedAt(*t)
	}
	return ftc
}

// SetUpdatedAt sets the "updated_at" field.
func (ftc *FileTagCreate) SetUpdatedAt(t time.Time) *FileTagCreate {
	ftc.mutation.SetUpdatedAt(t)
	return ftc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ftc *FileTagCreate) SetNillableUpdatedAt(t *time.Time) *FileTagCreate {
	if t != nil {
		ftc.SetUpdatedAt(*t)
	}
	return ftc
}

// SetStatus sets the "status" field.
func (ftc *FileTagCreate) SetStatus(u uint8) *FileTagCreate {
	ftc.mutation.SetStatus(u)
	return ftc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ftc *FileTagCreate) SetNillableStatus(u *uint8) *FileTagCreate {
	if u != nil {
		ftc.SetStatus(*u)
	}
	return ftc
}

// SetName sets the "name" field.
func (ftc *FileTagCreate) SetName(s string) *FileTagCreate {
	ftc.mutation.SetName(s)
	return ftc
}

// SetRemark sets the "remark" field.
func (ftc *FileTagCreate) SetRemark(s string) *FileTagCreate {
	ftc.mutation.SetRemark(s)
	return ftc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (ftc *FileTagCreate) SetNillableRemark(s *string) *FileTagCreate {
	if s != nil {
		ftc.SetRemark(*s)
	}
	return ftc
}

// SetID sets the "id" field.
func (ftc *FileTagCreate) SetID(u uint64) *FileTagCreate {
	ftc.mutation.SetID(u)
	return ftc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftc *FileTagCreate) AddFileIDs(ids ...uuid.UUID) *FileTagCreate {
	ftc.mutation.AddFileIDs(ids...)
	return ftc
}

// AddFiles adds the "files" edges to the File entity.
func (ftc *FileTagCreate) AddFiles(f ...*File) *FileTagCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Mutation returns the FileTagMutation object of the builder.
func (ftc *FileTagCreate) Mutation() *FileTagMutation {
	return ftc.mutation
}

// Save creates the FileTag in the database.
func (ftc *FileTagCreate) Save(ctx context.Context) (*FileTag, error) {
	ftc.defaults()
	return withHooks(ctx, ftc.sqlSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTagCreate) SaveX(ctx context.Context) *FileTag {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FileTagCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FileTagCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FileTagCreate) defaults() {
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		v := filetag.DefaultCreatedAt()
		ftc.mutation.SetCreatedAt(v)
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		v := filetag.DefaultUpdatedAt()
		ftc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ftc.mutation.Status(); !ok {
		v := filetag.DefaultStatus
		ftc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FileTagCreate) check() error {
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FileTag.created_at"`)}
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FileTag.updated_at"`)}
	}
	if _, ok := ftc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FileTag.name"`)}
	}
	return nil
}

func (ftc *FileTagCreate) sqlSave(ctx context.Context) (*FileTag, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	ftc.mutation.id = &_node.ID
	ftc.mutation.done = true
	return _node, nil
}

func (ftc *FileTagCreate) createSpec() (*FileTag, *sqlgraph.CreateSpec) {
	var (
		_node = &FileTag{config: ftc.config}
		_spec = sqlgraph.NewCreateSpec(filetag.Table, sqlgraph.NewFieldSpec(filetag.FieldID, field.TypeUint64))
	)
	if id, ok := ftc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ftc.mutation.CreatedAt(); ok {
		_spec.SetField(filetag.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ftc.mutation.UpdatedAt(); ok {
		_spec.SetField(filetag.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ftc.mutation.Status(); ok {
		_spec.SetField(filetag.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := ftc.mutation.Name(); ok {
		_spec.SetField(filetag.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ftc.mutation.Remark(); ok {
		_spec.SetField(filetag.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if nodes := ftc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   filetag.FilesTable,
			Columns: filetag.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileTagCreateBulk is the builder for creating many FileTag entities in bulk.
type FileTagCreateBulk struct {
	config
	err      error
	builders []*FileTagCreate
}

// Save creates the FileTag entities in the database.
func (ftcb *FileTagCreateBulk) Save(ctx context.Context) ([]*FileTag, error) {
	if ftcb.err != nil {
		return nil, ftcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FileTag, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileTagMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FileTagCreateBulk) SaveX(ctx context.Context) []*FileTag {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcb *FileTagCreateBulk) Exec(ctx context.Context) error {
	_, err := ftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcb *FileTagCreateBulk) ExecX(ctx context.Context) {
	if err := ftcb.Exec(ctx); err != nil {
		panic(err)
	}
}
