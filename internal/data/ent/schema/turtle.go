package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Turtle holds the schema definition for the Turtle entity.
type Turtle struct {
	ent.Schema
}

// Fields of the Post.
func (Turtle) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("qid"),
		field.String("title"),
		field.String("content"),
		field.String("answer"),
		field.Int32("category"),
		field.String("creator"),
		field.Int32("difficulty"),
		field.Int32("state"),
		field.Time("ctime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("mtime").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Post.
func (Turtle) Edges() []ent.Edge {
	return nil
}
