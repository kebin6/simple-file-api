// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/kebin6/simple-file-api/ent/file"
	"github.com/kebin6/simple-file-api/ent/filetag"
	"github.com/kebin6/simple-file-api/ent/predicate"
)

// FileTagQuery is the builder for querying FileTag entities.
type FileTagQuery struct {
	config
	ctx        *QueryContext
	order      []filetag.OrderOption
	inters     []Interceptor
	predicates []predicate.FileTag
	withFiles  *FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileTagQuery builder.
func (ftq *FileTagQuery) Where(ps ...predicate.FileTag) *FileTagQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit the number of records to be returned by this query.
func (ftq *FileTagQuery) Limit(limit int) *FileTagQuery {
	ftq.ctx.Limit = &limit
	return ftq
}

// Offset to start from.
func (ftq *FileTagQuery) Offset(offset int) *FileTagQuery {
	ftq.ctx.Offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FileTagQuery) Unique(unique bool) *FileTagQuery {
	ftq.ctx.Unique = &unique
	return ftq
}

// Order specifies how the records should be ordered.
func (ftq *FileTagQuery) Order(o ...filetag.OrderOption) *FileTagQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// QueryFiles chains the current query on the "files" edge.
func (ftq *FileTagQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: ftq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ftq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ftq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filetag.Table, filetag.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, filetag.FilesTable, filetag.FilesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ftq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileTag entity from the query.
// Returns a *NotFoundError when no FileTag was found.
func (ftq *FileTagQuery) First(ctx context.Context) (*FileTag, error) {
	nodes, err := ftq.Limit(1).All(setContextOp(ctx, ftq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filetag.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FileTagQuery) FirstX(ctx context.Context) *FileTag {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileTag ID from the query.
// Returns a *NotFoundError when no FileTag ID was found.
func (ftq *FileTagQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ftq.Limit(1).IDs(setContextOp(ctx, ftq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filetag.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FileTagQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileTag entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileTag entity is found.
// Returns a *NotFoundError when no FileTag entities are found.
func (ftq *FileTagQuery) Only(ctx context.Context) (*FileTag, error) {
	nodes, err := ftq.Limit(2).All(setContextOp(ctx, ftq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filetag.Label}
	default:
		return nil, &NotSingularError{filetag.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FileTagQuery) OnlyX(ctx context.Context) *FileTag {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileTag ID in the query.
// Returns a *NotSingularError when more than one FileTag ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FileTagQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ftq.Limit(2).IDs(setContextOp(ctx, ftq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filetag.Label}
	default:
		err = &NotSingularError{filetag.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FileTagQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileTags.
func (ftq *FileTagQuery) All(ctx context.Context) ([]*FileTag, error) {
	ctx = setContextOp(ctx, ftq.ctx, "All")
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FileTag, *FileTagQuery]()
	return withInterceptors[[]*FileTag](ctx, ftq, qr, ftq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FileTagQuery) AllX(ctx context.Context) []*FileTag {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileTag IDs.
func (ftq *FileTagQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if ftq.ctx.Unique == nil && ftq.path != nil {
		ftq.Unique(true)
	}
	ctx = setContextOp(ctx, ftq.ctx, "IDs")
	if err = ftq.Select(filetag.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FileTagQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FileTagQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Count")
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ftq, querierCount[*FileTagQuery](), ftq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FileTagQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FileTagQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Exist")
	switch _, err := ftq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FileTagQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileTagQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FileTagQuery) Clone() *FileTagQuery {
	if ftq == nil {
		return nil
	}
	return &FileTagQuery{
		config:     ftq.config,
		ctx:        ftq.ctx.Clone(),
		order:      append([]filetag.OrderOption{}, ftq.order...),
		inters:     append([]Interceptor{}, ftq.inters...),
		predicates: append([]predicate.FileTag{}, ftq.predicates...),
		withFiles:  ftq.withFiles.Clone(),
		// clone intermediate query.
		sql:  ftq.sql.Clone(),
		path: ftq.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (ftq *FileTagQuery) WithFiles(opts ...func(*FileQuery)) *FileTagQuery {
	query := (&FileClient{config: ftq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ftq.withFiles = query
	return ftq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileTag.Query().
//		GroupBy(filetag.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ftq *FileTagQuery) GroupBy(field string, fields ...string) *FileTagGroupBy {
	ftq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FileTagGroupBy{build: ftq}
	grbuild.flds = &ftq.ctx.Fields
	grbuild.label = filetag.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.FileTag.Query().
//		Select(filetag.FieldCreatedAt).
//		Scan(ctx, &v)
func (ftq *FileTagQuery) Select(fields ...string) *FileTagSelect {
	ftq.ctx.Fields = append(ftq.ctx.Fields, fields...)
	sbuild := &FileTagSelect{FileTagQuery: ftq}
	sbuild.label = filetag.Label
	sbuild.flds, sbuild.scan = &ftq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FileTagSelect configured with the given aggregations.
func (ftq *FileTagQuery) Aggregate(fns ...AggregateFunc) *FileTagSelect {
	return ftq.Select().Aggregate(fns...)
}

func (ftq *FileTagQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ftq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ftq); err != nil {
				return err
			}
		}
	}
	for _, f := range ftq.ctx.Fields {
		if !filetag.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.sql = prev
	}
	return nil
}

func (ftq *FileTagQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FileTag, error) {
	var (
		nodes       = []*FileTag{}
		_spec       = ftq.querySpec()
		loadedTypes = [1]bool{
			ftq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FileTag).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FileTag{config: ftq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ftq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ftq.withFiles; query != nil {
		if err := ftq.loadFiles(ctx, query, nodes,
			func(n *FileTag) { n.Edges.Files = []*File{} },
			func(n *FileTag, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ftq *FileTagQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*FileTag, init func(*FileTag), assign func(*FileTag, *File)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint64]*FileTag)
	nids := make(map[uuid.UUID]map[*FileTag]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(filetag.FilesTable)
		s.Join(joinT).On(s.C(file.FieldID), joinT.C(filetag.FilesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(filetag.FilesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(filetag.FilesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint64(values[0].(*sql.NullInt64).Int64)
				inValue := *values[1].(*uuid.UUID)
				if nids[inValue] == nil {
					nids[inValue] = map[*FileTag]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*File](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "files" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (ftq *FileTagQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftq.querySpec()
	_spec.Node.Columns = ftq.ctx.Fields
	if len(ftq.ctx.Fields) > 0 {
		_spec.Unique = ftq.ctx.Unique != nil && *ftq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ftq.driver, _spec)
}

func (ftq *FileTagQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(filetag.Table, filetag.Columns, sqlgraph.NewFieldSpec(filetag.FieldID, field.TypeUint64))
	_spec.From = ftq.sql
	if unique := ftq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ftq.path != nil {
		_spec.Unique = true
	}
	if fields := ftq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetag.FieldID)
		for i := range fields {
			if fields[i] != filetag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ftq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ftq *FileTagQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ftq.driver.Dialect())
	t1 := builder.Table(filetag.Table)
	columns := ftq.ctx.Fields
	if len(columns) == 0 {
		columns = filetag.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ftq.sql != nil {
		selector = ftq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ftq.ctx.Unique != nil && *ftq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ftq.predicates {
		p(selector)
	}
	for _, p := range ftq.order {
		p(selector)
	}
	if offset := ftq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FileTagGroupBy is the group-by builder for FileTag entities.
type FileTagGroupBy struct {
	selector
	build *FileTagQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FileTagGroupBy) Aggregate(fns ...AggregateFunc) *FileTagGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the selector query and scans the result into the given value.
func (ftgb *FileTagGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftgb.build.ctx, "GroupBy")
	if err := ftgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileTagQuery, *FileTagGroupBy](ctx, ftgb.build, ftgb, ftgb.build.inters, v)
}

func (ftgb *FileTagGroupBy) sqlScan(ctx context.Context, root *FileTagQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ftgb.fns))
	for _, fn := range ftgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ftgb.flds)+len(ftgb.fns))
		for _, f := range *ftgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ftgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FileTagSelect is the builder for selecting fields of FileTag entities.
type FileTagSelect struct {
	*FileTagQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fts *FileTagSelect) Aggregate(fns ...AggregateFunc) *FileTagSelect {
	fts.fns = append(fts.fns, fns...)
	return fts
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FileTagSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fts.ctx, "Select")
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FileTagQuery, *FileTagSelect](ctx, fts.FileTagQuery, fts, fts.inters, v)
}

func (fts *FileTagSelect) sqlScan(ctx context.Context, root *FileTagQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fts.fns))
	for _, fn := range fts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
