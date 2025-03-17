package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Idiom holds the schema definition for the Idiom entity.
type Idiom struct {
	ent.Schema
}

// Fields of the Post.
func (Idiom) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("iid"),
		field.String("name"),
		field.String("image"),
		field.Int32("difficulty"),
		field.String("creator"),
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
func (Idiom) Edges() []ent.Edge {
	return nil
}
