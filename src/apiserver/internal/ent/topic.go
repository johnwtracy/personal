// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/johnwtracy/personal/src/apiserver/internal/ent/topic"
)

// Topic is the model entity for the Topic schema.
type Topic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TopicQuery when eager-loading is set.
	Edges TopicEdges `json:"edges"`
}

// TopicEdges holds the relations/edges for other nodes in the graph.
type TopicEdges struct {
	// BlogPosts holds the value of the blog_posts edge.
	BlogPosts []*BlogPost
	// Projects holds the value of the projects edge.
	Projects []*Project
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BlogPostsOrErr returns the BlogPosts value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) BlogPostsOrErr() ([]*BlogPost, error) {
	if e.loadedTypes[0] {
		return e.BlogPosts, nil
	}
	return nil, &NotLoadedError{edge: "blog_posts"}
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e TopicEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[1] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Topic) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // tag
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Topic fields.
func (t *Topic) assignValues(values ...interface{}) error {
	if m, n := len(values), len(topic.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tag", values[0])
	} else if value.Valid {
		t.Tag = value.String
	}
	return nil
}

// QueryBlogPosts queries the blog_posts edge of the Topic.
func (t *Topic) QueryBlogPosts() *BlogPostQuery {
	return (&TopicClient{config: t.config}).QueryBlogPosts(t)
}

// QueryProjects queries the projects edge of the Topic.
func (t *Topic) QueryProjects() *ProjectQuery {
	return (&TopicClient{config: t.config}).QueryProjects(t)
}

// Update returns a builder for updating this Topic.
// Note that, you need to call Topic.Unwrap() before calling this method, if this Topic
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Topic) Update() *TopicUpdateOne {
	return (&TopicClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Topic) Unwrap() *Topic {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Topic is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Topic) String() string {
	var builder strings.Builder
	builder.WriteString("Topic(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", tag=")
	builder.WriteString(t.Tag)
	builder.WriteByte(')')
	return builder.String()
}

// Topics is a parsable slice of Topic.
type Topics []*Topic

func (t Topics) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}