// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"github.com/johnwtracy/personal/src/apiserver/internal/project/ent/blogpost"
	"github.com/johnwtracy/personal/src/apiserver/internal/project/ent/topic"
)

// BlogPostCreate is the builder for creating a BlogPost entity.
type BlogPostCreate struct {
	config
	mutation *BlogPostMutation
	hooks    []Hook
}

// SetHead sets the head field.
func (bpc *BlogPostCreate) SetHead(s string) *BlogPostCreate {
	bpc.mutation.SetHead(s)
	return bpc
}

// SetBody sets the body field.
func (bpc *BlogPostCreate) SetBody(s string) *BlogPostCreate {
	bpc.mutation.SetBody(s)
	return bpc
}

// SetCreateTime sets the create_time field.
func (bpc *BlogPostCreate) SetCreateTime(t time.Time) *BlogPostCreate {
	bpc.mutation.SetCreateTime(t)
	return bpc
}

// SetNillableCreateTime sets the create_time field if the given value is not nil.
func (bpc *BlogPostCreate) SetNillableCreateTime(t *time.Time) *BlogPostCreate {
	if t != nil {
		bpc.SetCreateTime(*t)
	}
	return bpc
}

// SetUpdateTime sets the update_time field.
func (bpc *BlogPostCreate) SetUpdateTime(t time.Time) *BlogPostCreate {
	bpc.mutation.SetUpdateTime(t)
	return bpc
}

// SetNillableUpdateTime sets the update_time field if the given value is not nil.
func (bpc *BlogPostCreate) SetNillableUpdateTime(t *time.Time) *BlogPostCreate {
	if t != nil {
		bpc.SetUpdateTime(*t)
	}
	return bpc
}

// SetID sets the id field.
func (bpc *BlogPostCreate) SetID(u uuid.UUID) *BlogPostCreate {
	bpc.mutation.SetID(u)
	return bpc
}

// AddTagIDs adds the tags edge to Topic by ids.
func (bpc *BlogPostCreate) AddTagIDs(ids ...int) *BlogPostCreate {
	bpc.mutation.AddTagIDs(ids...)
	return bpc
}

// AddTags adds the tags edges to Topic.
func (bpc *BlogPostCreate) AddTags(t ...*Topic) *BlogPostCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bpc.AddTagIDs(ids...)
}

// Mutation returns the BlogPostMutation object of the builder.
func (bpc *BlogPostCreate) Mutation() *BlogPostMutation {
	return bpc.mutation
}

// Save creates the BlogPost in the database.
func (bpc *BlogPostCreate) Save(ctx context.Context) (*BlogPost, error) {
	if err := bpc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *BlogPost
	)
	if len(bpc.hooks) == 0 {
		node, err = bpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlogPostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bpc.mutation = mutation
			node, err = bpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bpc.hooks) - 1; i >= 0; i-- {
			mut = bpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bpc *BlogPostCreate) SaveX(ctx context.Context) *BlogPost {
	v, err := bpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bpc *BlogPostCreate) preSave() error {
	if _, ok := bpc.mutation.Head(); !ok {
		return &ValidationError{Name: "head", err: errors.New("ent: missing required field \"head\"")}
	}
	if v, ok := bpc.mutation.Head(); ok {
		if err := blogpost.HeadValidator(v); err != nil {
			return &ValidationError{Name: "head", err: fmt.Errorf("ent: validator failed for field \"head\": %w", err)}
		}
	}
	if _, ok := bpc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New("ent: missing required field \"body\"")}
	}
	if v, ok := bpc.mutation.Body(); ok {
		if err := blogpost.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf("ent: validator failed for field \"body\": %w", err)}
		}
	}
	if _, ok := bpc.mutation.CreateTime(); !ok {
		v := blogpost.DefaultCreateTime()
		bpc.mutation.SetCreateTime(v)
	}
	if _, ok := bpc.mutation.UpdateTime(); !ok {
		v := blogpost.DefaultUpdateTime()
		bpc.mutation.SetUpdateTime(v)
	}
	return nil
}

func (bpc *BlogPostCreate) sqlSave(ctx context.Context) (*BlogPost, error) {
	bp, _spec := bpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return bp, nil
}

func (bpc *BlogPostCreate) createSpec() (*BlogPost, *sqlgraph.CreateSpec) {
	var (
		bp    = &BlogPost{config: bpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blogpost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blogpost.FieldID,
			},
		}
	)
	if id, ok := bpc.mutation.ID(); ok {
		bp.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bpc.mutation.Head(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blogpost.FieldHead,
		})
		bp.Head = value
	}
	if value, ok := bpc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: blogpost.FieldBody,
		})
		bp.Body = value
	}
	if value, ok := bpc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blogpost.FieldCreateTime,
		})
		bp.CreateTime = value
	}
	if value, ok := bpc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: blogpost.FieldUpdateTime,
		})
		bp.UpdateTime = value
	}
	if nodes := bpc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   blogpost.TagsTable,
			Columns: blogpost.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: topic.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return bp, _spec
}

// BlogPostCreateBulk is the builder for creating a bulk of BlogPost entities.
type BlogPostCreateBulk struct {
	config
	builders []*BlogPostCreate
}

// Save creates the BlogPost entities in the database.
func (bpcb *BlogPostCreateBulk) Save(ctx context.Context) ([]*BlogPost, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bpcb.builders))
	nodes := make([]*BlogPost, len(bpcb.builders))
	mutators := make([]Mutator, len(bpcb.builders))
	for i := range bpcb.builders {
		func(i int, root context.Context) {
			builder := bpcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*BlogPostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (bpcb *BlogPostCreateBulk) SaveX(ctx context.Context) []*BlogPost {
	v, err := bpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
