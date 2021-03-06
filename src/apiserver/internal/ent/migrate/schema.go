// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// BlogPostsColumns holds the columns for the "blog_posts" table.
	BlogPostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "head", Type: field.TypeString, Size: 32},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
	}
	// BlogPostsTable holds the schema information for the "blog_posts" table.
	BlogPostsTable = &schema.Table{
		Name:        "blog_posts",
		Columns:     BlogPostsColumns,
		PrimaryKey:  []*schema.Column{BlogPostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "head", Type: field.TypeString, Size: 32},
		{Name: "body", Type: field.TypeString, Size: 2147483647},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "started", Type: field.TypeTime},
		{Name: "completed", Type: field.TypeTime, Nullable: true},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:        "projects",
		Columns:     ProjectsColumns,
		PrimaryKey:  []*schema.Column{ProjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TopicsColumns holds the columns for the "topics" table.
	TopicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag", Type: field.TypeString, Unique: true},
	}
	// TopicsTable holds the schema information for the "topics" table.
	TopicsTable = &schema.Table{
		Name:        "topics",
		Columns:     TopicsColumns,
		PrimaryKey:  []*schema.Column{TopicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// TopicBlogPostsColumns holds the columns for the "topic_blog_posts" table.
	TopicBlogPostsColumns = []*schema.Column{
		{Name: "topic_id", Type: field.TypeInt},
		{Name: "blog_post_id", Type: field.TypeInt},
	}
	// TopicBlogPostsTable holds the schema information for the "topic_blog_posts" table.
	TopicBlogPostsTable = &schema.Table{
		Name:       "topic_blog_posts",
		Columns:    TopicBlogPostsColumns,
		PrimaryKey: []*schema.Column{TopicBlogPostsColumns[0], TopicBlogPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "topic_blog_posts_topic_id",
				Columns: []*schema.Column{TopicBlogPostsColumns[0]},

				RefColumns: []*schema.Column{TopicsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "topic_blog_posts_blog_post_id",
				Columns: []*schema.Column{TopicBlogPostsColumns[1]},

				RefColumns: []*schema.Column{BlogPostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TopicProjectsColumns holds the columns for the "topic_projects" table.
	TopicProjectsColumns = []*schema.Column{
		{Name: "topic_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// TopicProjectsTable holds the schema information for the "topic_projects" table.
	TopicProjectsTable = &schema.Table{
		Name:       "topic_projects",
		Columns:    TopicProjectsColumns,
		PrimaryKey: []*schema.Column{TopicProjectsColumns[0], TopicProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "topic_projects_topic_id",
				Columns: []*schema.Column{TopicProjectsColumns[0]},

				RefColumns: []*schema.Column{TopicsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:  "topic_projects_project_id",
				Columns: []*schema.Column{TopicProjectsColumns[1]},

				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlogPostsTable,
		ProjectsTable,
		TopicsTable,
		TopicBlogPostsTable,
		TopicProjectsTable,
	}
)

func init() {
	TopicBlogPostsTable.ForeignKeys[0].RefTable = TopicsTable
	TopicBlogPostsTable.ForeignKeys[1].RefTable = BlogPostsTable
	TopicProjectsTable.ForeignKeys[0].RefTable = TopicsTable
	TopicProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
}
