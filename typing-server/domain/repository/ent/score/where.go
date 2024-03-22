// Code generated by ent, DO NOT EDIT.

package score

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/su-its/typing/typing-server/domain/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldID, id))
}

// Keystrokes applies equality check predicate on the "keystrokes" field. It's identical to KeystrokesEQ.
func Keystrokes(v int) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldKeystrokes, v))
}

// Accuracy applies equality check predicate on the "accuracy" field. It's identical to AccuracyEQ.
func Accuracy(v float64) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldAccuracy, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldCreatedAt, v))
}

// KeystrokesEQ applies the EQ predicate on the "keystrokes" field.
func KeystrokesEQ(v int) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldKeystrokes, v))
}

// KeystrokesNEQ applies the NEQ predicate on the "keystrokes" field.
func KeystrokesNEQ(v int) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldKeystrokes, v))
}

// KeystrokesIn applies the In predicate on the "keystrokes" field.
func KeystrokesIn(vs ...int) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldKeystrokes, vs...))
}

// KeystrokesNotIn applies the NotIn predicate on the "keystrokes" field.
func KeystrokesNotIn(vs ...int) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldKeystrokes, vs...))
}

// KeystrokesGT applies the GT predicate on the "keystrokes" field.
func KeystrokesGT(v int) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldKeystrokes, v))
}

// KeystrokesGTE applies the GTE predicate on the "keystrokes" field.
func KeystrokesGTE(v int) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldKeystrokes, v))
}

// KeystrokesLT applies the LT predicate on the "keystrokes" field.
func KeystrokesLT(v int) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldKeystrokes, v))
}

// KeystrokesLTE applies the LTE predicate on the "keystrokes" field.
func KeystrokesLTE(v int) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldKeystrokes, v))
}

// AccuracyEQ applies the EQ predicate on the "accuracy" field.
func AccuracyEQ(v float64) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldAccuracy, v))
}

// AccuracyNEQ applies the NEQ predicate on the "accuracy" field.
func AccuracyNEQ(v float64) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldAccuracy, v))
}

// AccuracyIn applies the In predicate on the "accuracy" field.
func AccuracyIn(vs ...float64) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldAccuracy, vs...))
}

// AccuracyNotIn applies the NotIn predicate on the "accuracy" field.
func AccuracyNotIn(vs ...float64) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldAccuracy, vs...))
}

// AccuracyGT applies the GT predicate on the "accuracy" field.
func AccuracyGT(v float64) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldAccuracy, v))
}

// AccuracyGTE applies the GTE predicate on the "accuracy" field.
func AccuracyGTE(v float64) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldAccuracy, v))
}

// AccuracyLT applies the LT predicate on the "accuracy" field.
func AccuracyLT(v float64) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldAccuracy, v))
}

// AccuracyLTE applies the LTE predicate on the "accuracy" field.
func AccuracyLTE(v float64) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldAccuracy, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Score) predicate.Score {
	return predicate.Score(sql.NotPredicates(p))
}
