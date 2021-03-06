// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/blogpost"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/predicate"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/topic"
)

// BlogPostQuery is the builder for querying BlogPost entities.
type BlogPostQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.BlogPost
	// eager-loading edges.
	withTags *TopicQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (bpq *BlogPostQuery) Where(ps ...predicate.BlogPost) *BlogPostQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit adds a limit step to the query.
func (bpq *BlogPostQuery) Limit(limit int) *BlogPostQuery {
	bpq.limit = &limit
	return bpq
}

// Offset adds an offset step to the query.
func (bpq *BlogPostQuery) Offset(offset int) *BlogPostQuery {
	bpq.offset = &offset
	return bpq
}

// Order adds an order step to the query.
func (bpq *BlogPostQuery) Order(o ...OrderFunc) *BlogPostQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryTags chains the current query on the tags edge.
func (bpq *BlogPostQuery) QueryTags() *TopicQuery {
	query := &TopicQuery{config: bpq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blogpost.Table, blogpost.FieldID, bpq.sqlQuery()),
			sqlgraph.To(topic.Table, topic.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, blogpost.TagsTable, blogpost.TagsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BlogPost entity in the query. Returns *NotFoundError when no blogpost was found.
func (bpq *BlogPostQuery) First(ctx context.Context) (*BlogPost, error) {
	bps, err := bpq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(bps) == 0 {
		return nil, &NotFoundError{blogpost.Label}
	}
	return bps[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstX(ctx context.Context) *BlogPost {
	bp, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return bp
}

// FirstID returns the first BlogPost id in the query. Returns *NotFoundError when no id was found.
func (bpq *BlogPostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blogpost.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (bpq *BlogPostQuery) FirstXID(ctx context.Context) int {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only BlogPost entity in the query, returns an error if not exactly one entity was returned.
func (bpq *BlogPostQuery) Only(ctx context.Context) (*BlogPost, error) {
	bps, err := bpq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(bps) {
	case 1:
		return bps[0], nil
	case 0:
		return nil, &NotFoundError{blogpost.Label}
	default:
		return nil, &NotSingularError{blogpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyX(ctx context.Context) *BlogPost {
	bp, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return bp
}

// OnlyID returns the only BlogPost id in the query, returns an error if not exactly one id was returned.
func (bpq *BlogPostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bpq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = &NotSingularError{blogpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BlogPostQuery) OnlyIDX(ctx context.Context) int {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlogPosts.
func (bpq *BlogPostQuery) All(ctx context.Context) ([]*BlogPost, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bpq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BlogPostQuery) AllX(ctx context.Context) []*BlogPost {
	bps, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return bps
}

// IDs executes the query and returns a list of BlogPost ids.
func (bpq *BlogPostQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bpq.Select(blogpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BlogPostQuery) IDsX(ctx context.Context) []int {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BlogPostQuery) Count(ctx context.Context) (int, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bpq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BlogPostQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BlogPostQuery) Exist(ctx context.Context) (bool, error) {
	if err := bpq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bpq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BlogPostQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BlogPostQuery) Clone() *BlogPostQuery {
	return &BlogPostQuery{
		config:     bpq.config,
		limit:      bpq.limit,
		offset:     bpq.offset,
		order:      append([]OrderFunc{}, bpq.order...),
		unique:     append([]string{}, bpq.unique...),
		predicates: append([]predicate.BlogPost{}, bpq.predicates...),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

//  WithTags tells the query-builder to eager-loads the nodes that are connected to
// the "tags" edge. The optional arguments used to configure the query builder of the edge.
func (bpq *BlogPostQuery) WithTags(opts ...func(*TopicQuery)) *BlogPostQuery {
	query := &TopicQuery{config: bpq.config}
	for _, opt := range opts {
		opt(query)
	}
	bpq.withTags = query
	return bpq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Head string `json:"head,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		GroupBy(blogpost.FieldHead).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bpq *BlogPostQuery) GroupBy(field string, fields ...string) *BlogPostGroupBy {
	group := &BlogPostGroupBy{config: bpq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bpq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Head string `json:"head,omitempty"`
//	}
//
//	client.BlogPost.Query().
//		Select(blogpost.FieldHead).
//		Scan(ctx, &v)
//
func (bpq *BlogPostQuery) Select(field string, fields ...string) *BlogPostSelect {
	selector := &BlogPostSelect{config: bpq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bpq.sqlQuery(), nil
	}
	return selector
}

func (bpq *BlogPostQuery) prepareQuery(ctx context.Context) error {
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BlogPostQuery) sqlAll(ctx context.Context) ([]*BlogPost, error) {
	var (
		nodes       = []*BlogPost{}
		_spec       = bpq.querySpec()
		loadedTypes = [1]bool{
			bpq.withTags != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &BlogPost{config: bpq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bpq.withTags; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*BlogPost, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
		}
		var (
			edgeids []int
			edges   = make(map[int][]*BlogPost)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: true,
				Table:   blogpost.TagsTable,
				Columns: blogpost.TagsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(blogpost.TagsPrimaryKey[1], fks...))
			},

			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				edgeids = append(edgeids, inValue)
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, bpq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "tags": %v`, err)
		}
		query.Where(topic.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "tags" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Tags = append(nodes[i].Edges.Tags, n)
			}
		}
	}

	return nodes, nil
}

func (bpq *BlogPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BlogPostQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bpq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (bpq *BlogPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blogpost.Table,
			Columns: blogpost.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blogpost.FieldID,
			},
		},
		From:   bpq.sql,
		Unique: true,
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BlogPostQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(blogpost.Table)
	selector := builder.Select(t1.Columns(blogpost.Columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(blogpost.Columns...)...)
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlogPostGroupBy is the builder for group-by BlogPost entities.
type BlogPostGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BlogPostGroupBy) Aggregate(fns ...AggregateFunc) *BlogPostGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the group-by query and scan the result into the given value.
func (bpgb *BlogPostGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bpgb.path(ctx)
	if err != nil {
		return err
	}
	bpgb.sql = query
	return bpgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bpgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) StringsX(ctx context.Context) []string {
	v, err := bpgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bpgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) StringX(ctx context.Context) string {
	v, err := bpgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) IntsX(ctx context.Context) []int {
	v, err := bpgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bpgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) IntX(ctx context.Context) int {
	v, err := bpgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bpgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bpgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bpgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bpgb.fields) > 1 {
		return nil, errors.New("ent: BlogPostGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bpgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bpgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (bpgb *BlogPostGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bpgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bpgb *BlogPostGroupBy) BoolX(ctx context.Context) bool {
	v, err := bpgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bpgb *BlogPostGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bpgb.sqlQuery().Query()
	if err := bpgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bpgb *BlogPostGroupBy) sqlQuery() *sql.Selector {
	selector := bpgb.sql
	columns := make([]string, 0, len(bpgb.fields)+len(bpgb.fns))
	columns = append(columns, bpgb.fields...)
	for _, fn := range bpgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(bpgb.fields...)
}

// BlogPostSelect is the builder for select fields of BlogPost entities.
type BlogPostSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (bps *BlogPostSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := bps.path(ctx)
	if err != nil {
		return err
	}
	bps.sql = query
	return bps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bps *BlogPostSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bps *BlogPostSelect) StringsX(ctx context.Context) []string {
	v, err := bps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bps *BlogPostSelect) StringX(ctx context.Context) string {
	v, err := bps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bps *BlogPostSelect) IntsX(ctx context.Context) []int {
	v, err := bps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bps *BlogPostSelect) IntX(ctx context.Context) int {
	v, err := bps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bps *BlogPostSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bps *BlogPostSelect) Float64X(ctx context.Context) float64 {
	v, err := bps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bps.fields) > 1 {
		return nil, errors.New("ent: BlogPostSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bps *BlogPostSelect) BoolsX(ctx context.Context) []bool {
	v, err := bps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (bps *BlogPostSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{blogpost.Label}
	default:
		err = fmt.Errorf("ent: BlogPostSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bps *BlogPostSelect) BoolX(ctx context.Context) bool {
	v, err := bps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bps *BlogPostSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bps.sqlQuery().Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bps *BlogPostSelect) sqlQuery() sql.Querier {
	selector := bps.sql
	selector.Select(selector.Columns(bps.fields...)...)
	return selector
}
