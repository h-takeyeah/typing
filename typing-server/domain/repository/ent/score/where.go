// Code generated by ent, DO NOT EDIT.

package score

import (
	"ent/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
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

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float64) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldScore, v))
}

// StartedAt applies equality check predicate on the "startedAt" field. It's identical to StartedAtEQ.
func StartedAt(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldStartedAt, v))
}

// EndedAt applies equality check predicate on the "endedAt" field. It's identical to EndedAtEQ.
func EndedAt(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldEndedAt, v))
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

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float64) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float64) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float64) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float64) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float64) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float64) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float64) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float64) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldScore, v))
}

// StartedAtEQ applies the EQ predicate on the "startedAt" field.
func StartedAtEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldStartedAt, v))
}

// StartedAtNEQ applies the NEQ predicate on the "startedAt" field.
func StartedAtNEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldStartedAt, v))
}

// StartedAtIn applies the In predicate on the "startedAt" field.
func StartedAtIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldStartedAt, vs...))
}

// StartedAtNotIn applies the NotIn predicate on the "startedAt" field.
func StartedAtNotIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldStartedAt, vs...))
}

// StartedAtGT applies the GT predicate on the "startedAt" field.
func StartedAtGT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldStartedAt, v))
}

// StartedAtGTE applies the GTE predicate on the "startedAt" field.
func StartedAtGTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldStartedAt, v))
}

// StartedAtLT applies the LT predicate on the "startedAt" field.
func StartedAtLT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldStartedAt, v))
}

// StartedAtLTE applies the LTE predicate on the "startedAt" field.
func StartedAtLTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldStartedAt, v))
}

// EndedAtEQ applies the EQ predicate on the "endedAt" field.
func EndedAtEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldEQ(FieldEndedAt, v))
}

// EndedAtNEQ applies the NEQ predicate on the "endedAt" field.
func EndedAtNEQ(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldNEQ(FieldEndedAt, v))
}

// EndedAtIn applies the In predicate on the "endedAt" field.
func EndedAtIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldIn(FieldEndedAt, vs...))
}

// EndedAtNotIn applies the NotIn predicate on the "endedAt" field.
func EndedAtNotIn(vs ...time.Time) predicate.Score {
	return predicate.Score(sql.FieldNotIn(FieldEndedAt, vs...))
}

// EndedAtGT applies the GT predicate on the "endedAt" field.
func EndedAtGT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGT(FieldEndedAt, v))
}

// EndedAtGTE applies the GTE predicate on the "endedAt" field.
func EndedAtGTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldGTE(FieldEndedAt, v))
}

// EndedAtLT applies the LT predicate on the "endedAt" field.
func EndedAtLT(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLT(FieldEndedAt, v))
}

// EndedAtLTE applies the LTE predicate on the "endedAt" field.
func EndedAtLTE(v time.Time) predicate.Score {
	return predicate.Score(sql.FieldLTE(FieldEndedAt, v))
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
