// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/blogpost"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/predicate"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/project"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/topic"
)

// TopicUpdate is the builder for updating Topic entities.
type TopicUpdate struct {
	config
	hooks      []Hook
	mutation   *TopicMutation
	predicates []predicate.Topic
}

// Where adds a new predicate for the builder.
func (tu *TopicUpdate) Where(ps ...predicate.Topic) *TopicUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetTag sets the tag field.
func (tu *TopicUpdate) SetTag(s string) *TopicUpdate {
	tu.mutation.SetTag(s)
	return tu
}

// AddBlogPostIDs adds the blog_posts edge to BlogPost by ids.
func (tu *TopicUpdate) AddBlogPostIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddBlogPostIDs(ids...)
	return tu
}

// AddBlogPosts adds the blog_posts edges to BlogPost.
func (tu *TopicUpdate) AddBlogPosts(b ...*BlogPost) *TopicUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.AddBlogPostIDs(ids...)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (tu *TopicUpdate) AddProjectIDs(ids ...int) *TopicUpdate {
	tu.mutation.AddProjectIDs(ids...)
	return tu
}

// AddProjects adds the projects edges to Project.
func (tu *TopicUpdate) AddProjects(p ...*Project) *TopicUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddProjectIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tu *TopicUpdate) Mutation() *TopicMutation {
	return tu.mutation
}

// RemoveBlogPostIDs removes the blog_posts edge to BlogPost by ids.
func (tu *TopicUpdate) RemoveBlogPostIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveBlogPostIDs(ids...)
	return tu
}

// RemoveBlogPosts removes blog_posts edges to BlogPost.
func (tu *TopicUpdate) RemoveBlogPosts(b ...*BlogPost) *TopicUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.RemoveBlogPostIDs(ids...)
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (tu *TopicUpdate) RemoveProjectIDs(ids ...int) *TopicUpdate {
	tu.mutation.RemoveProjectIDs(ids...)
	return tu
}

// RemoveProjects removes projects edges to Project.
func (tu *TopicUpdate) RemoveProjects(p ...*Project) *TopicUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TopicUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := tu.mutation.Tag(); ok {
		if err := topic.TagValidator(v); err != nil {
			return 0, &ValidationError{Name: "tag", err: fmt.Errorf("ent: validator failed for field \"tag\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TopicUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TopicUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TopicUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TopicUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldTag,
		})
	}
	if nodes := tu.mutation.RemovedBlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.BlogPostsTable,
			Columns: topic.BlogPostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.BlogPostsTable,
			Columns: topic.BlogPostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tu.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: topic.ProjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: topic.ProjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TopicUpdateOne is the builder for updating a single Topic entity.
type TopicUpdateOne struct {
	config
	hooks    []Hook
	mutation *TopicMutation
}

// SetTag sets the tag field.
func (tuo *TopicUpdateOne) SetTag(s string) *TopicUpdateOne {
	tuo.mutation.SetTag(s)
	return tuo
}

// AddBlogPostIDs adds the blog_posts edge to BlogPost by ids.
func (tuo *TopicUpdateOne) AddBlogPostIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddBlogPostIDs(ids...)
	return tuo
}

// AddBlogPosts adds the blog_posts edges to BlogPost.
func (tuo *TopicUpdateOne) AddBlogPosts(b ...*BlogPost) *TopicUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.AddBlogPostIDs(ids...)
}

// AddProjectIDs adds the projects edge to Project by ids.
func (tuo *TopicUpdateOne) AddProjectIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.AddProjectIDs(ids...)
	return tuo
}

// AddProjects adds the projects edges to Project.
func (tuo *TopicUpdateOne) AddProjects(p ...*Project) *TopicUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddProjectIDs(ids...)
}

// Mutation returns the TopicMutation object of the builder.
func (tuo *TopicUpdateOne) Mutation() *TopicMutation {
	return tuo.mutation
}

// RemoveBlogPostIDs removes the blog_posts edge to BlogPost by ids.
func (tuo *TopicUpdateOne) RemoveBlogPostIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveBlogPostIDs(ids...)
	return tuo
}

// RemoveBlogPosts removes blog_posts edges to BlogPost.
func (tuo *TopicUpdateOne) RemoveBlogPosts(b ...*BlogPost) *TopicUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.RemoveBlogPostIDs(ids...)
}

// RemoveProjectIDs removes the projects edge to Project by ids.
func (tuo *TopicUpdateOne) RemoveProjectIDs(ids ...int) *TopicUpdateOne {
	tuo.mutation.RemoveProjectIDs(ids...)
	return tuo
}

// RemoveProjects removes projects edges to Project.
func (tuo *TopicUpdateOne) RemoveProjects(p ...*Project) *TopicUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemoveProjectIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TopicUpdateOne) Save(ctx context.Context) (*Topic, error) {
	if v, ok := tuo.mutation.Tag(); ok {
		if err := topic.TagValidator(v); err != nil {
			return nil, &ValidationError{Name: "tag", err: fmt.Errorf("ent: validator failed for field \"tag\": %w", err)}
		}
	}

	var (
		err  error
		node *Topic
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TopicUpdateOne) SaveX(ctx context.Context) *Topic {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TopicUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TopicUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TopicUpdateOne) sqlSave(ctx context.Context) (t *Topic, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topic.Table,
			Columns: topic.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: topic.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Topic.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topic.FieldTag,
		})
	}
	if nodes := tuo.mutation.RemovedBlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.BlogPostsTable,
			Columns: topic.BlogPostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BlogPostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.BlogPostsTable,
			Columns: topic.BlogPostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: blogpost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := tuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: topic.ProjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   topic.ProjectsTable,
			Columns: topic.ProjectsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Topic{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topic.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
