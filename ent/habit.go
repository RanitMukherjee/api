// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/habit"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Habit is the model entity for the Habit schema.
type Habit struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt    time.Time `json:"created_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Habit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case habit.FieldID:
			values[i] = new(sql.NullInt64)
		case habit.FieldName, habit.FieldDescription:
			values[i] = new(sql.NullString)
		case habit.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Habit fields.
func (h *Habit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case habit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case habit.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		case habit.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				h.Description = value.String
			}
		case habit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Habit.
// This includes values selected through modifiers, order, etc.
func (h *Habit) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// Update returns a builder for updating this Habit.
// Note that you need to call Habit.Unwrap() before calling this method if this Habit
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Habit) Update() *HabitUpdateOne {
	return NewHabitClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Habit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Habit) Unwrap() *Habit {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Habit is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Habit) String() string {
	var builder strings.Builder
	builder.WriteString("Habit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(h.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Habits is a parsable slice of Habit.
type Habits []*Habit
