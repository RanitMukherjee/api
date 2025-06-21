package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "time"
)

type Habit struct {
    ent.Schema
}

func (Habit) Fields() []ent.Field {
    return []ent.Field{
        field.String("name").NotEmpty(),
        field.String("description").Optional(),
        field.Time("created_at").Default(time.Now),
    }
}
