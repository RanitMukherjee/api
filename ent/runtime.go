// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/habit"
	"api/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	habitFields := schema.Habit{}.Fields()
	_ = habitFields
	// habitDescName is the schema descriptor for name field.
	habitDescName := habitFields[0].Descriptor()
	// habit.NameValidator is a validator for the "name" field. It is called by the builders before save.
	habit.NameValidator = habitDescName.Validators[0].(func(string) error)
	// habitDescCreatedAt is the schema descriptor for created_at field.
	habitDescCreatedAt := habitFields[2].Descriptor()
	// habit.DefaultCreatedAt holds the default value on creation for the created_at field.
	habit.DefaultCreatedAt = habitDescCreatedAt.Default.(func() time.Time)
}
