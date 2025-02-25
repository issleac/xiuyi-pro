// Code generated by ent, DO NOT EDIT.

package turtle

import (
	"time"
	"xiuyiPro/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldID, id))
}

// Qid applies equality check predicate on the "qid" field. It's identical to QidEQ.
func Qid(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldQid, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldTitle, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldContent, v))
}

// Answer applies equality check predicate on the "answer" field. It's identical to AnswerEQ.
func Answer(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldAnswer, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCategory, v))
}

// Creator applies equality check predicate on the "creator" field. It's identical to CreatorEQ.
func Creator(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCreator, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldState, v))
}

// Ctime applies equality check predicate on the "ctime" field. It's identical to CtimeEQ.
func Ctime(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCtime, v))
}

// Mtime applies equality check predicate on the "mtime" field. It's identical to MtimeEQ.
func Mtime(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldMtime, v))
}

// QidEQ applies the EQ predicate on the "qid" field.
func QidEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldQid, v))
}

// QidNEQ applies the NEQ predicate on the "qid" field.
func QidNEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldQid, v))
}

// QidIn applies the In predicate on the "qid" field.
func QidIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldQid, vs...))
}

// QidNotIn applies the NotIn predicate on the "qid" field.
func QidNotIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldQid, vs...))
}

// QidGT applies the GT predicate on the "qid" field.
func QidGT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldQid, v))
}

// QidGTE applies the GTE predicate on the "qid" field.
func QidGTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldQid, v))
}

// QidLT applies the LT predicate on the "qid" field.
func QidLT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldQid, v))
}

// QidLTE applies the LTE predicate on the "qid" field.
func QidLTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldQid, v))
}

// QidContains applies the Contains predicate on the "qid" field.
func QidContains(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContains(FieldQid, v))
}

// QidHasPrefix applies the HasPrefix predicate on the "qid" field.
func QidHasPrefix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasPrefix(FieldQid, v))
}

// QidHasSuffix applies the HasSuffix predicate on the "qid" field.
func QidHasSuffix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasSuffix(FieldQid, v))
}

// QidEqualFold applies the EqualFold predicate on the "qid" field.
func QidEqualFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEqualFold(FieldQid, v))
}

// QidContainsFold applies the ContainsFold predicate on the "qid" field.
func QidContainsFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContainsFold(FieldQid, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContainsFold(FieldTitle, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContainsFold(FieldContent, v))
}

// AnswerEQ applies the EQ predicate on the "answer" field.
func AnswerEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldAnswer, v))
}

// AnswerNEQ applies the NEQ predicate on the "answer" field.
func AnswerNEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldAnswer, v))
}

// AnswerIn applies the In predicate on the "answer" field.
func AnswerIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldAnswer, vs...))
}

// AnswerNotIn applies the NotIn predicate on the "answer" field.
func AnswerNotIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldAnswer, vs...))
}

// AnswerGT applies the GT predicate on the "answer" field.
func AnswerGT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldAnswer, v))
}

// AnswerGTE applies the GTE predicate on the "answer" field.
func AnswerGTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldAnswer, v))
}

// AnswerLT applies the LT predicate on the "answer" field.
func AnswerLT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldAnswer, v))
}

// AnswerLTE applies the LTE predicate on the "answer" field.
func AnswerLTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldAnswer, v))
}

// AnswerContains applies the Contains predicate on the "answer" field.
func AnswerContains(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContains(FieldAnswer, v))
}

// AnswerHasPrefix applies the HasPrefix predicate on the "answer" field.
func AnswerHasPrefix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasPrefix(FieldAnswer, v))
}

// AnswerHasSuffix applies the HasSuffix predicate on the "answer" field.
func AnswerHasSuffix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasSuffix(FieldAnswer, v))
}

// AnswerEqualFold applies the EqualFold predicate on the "answer" field.
func AnswerEqualFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEqualFold(FieldAnswer, v))
}

// AnswerContainsFold applies the ContainsFold predicate on the "answer" field.
func AnswerContainsFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContainsFold(FieldAnswer, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldCategory, v))
}

// CreatorEQ applies the EQ predicate on the "creator" field.
func CreatorEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCreator, v))
}

// CreatorNEQ applies the NEQ predicate on the "creator" field.
func CreatorNEQ(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldCreator, v))
}

// CreatorIn applies the In predicate on the "creator" field.
func CreatorIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldCreator, vs...))
}

// CreatorNotIn applies the NotIn predicate on the "creator" field.
func CreatorNotIn(vs ...string) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldCreator, vs...))
}

// CreatorGT applies the GT predicate on the "creator" field.
func CreatorGT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldCreator, v))
}

// CreatorGTE applies the GTE predicate on the "creator" field.
func CreatorGTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldCreator, v))
}

// CreatorLT applies the LT predicate on the "creator" field.
func CreatorLT(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldCreator, v))
}

// CreatorLTE applies the LTE predicate on the "creator" field.
func CreatorLTE(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldCreator, v))
}

// CreatorContains applies the Contains predicate on the "creator" field.
func CreatorContains(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContains(FieldCreator, v))
}

// CreatorHasPrefix applies the HasPrefix predicate on the "creator" field.
func CreatorHasPrefix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasPrefix(FieldCreator, v))
}

// CreatorHasSuffix applies the HasSuffix predicate on the "creator" field.
func CreatorHasSuffix(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldHasSuffix(FieldCreator, v))
}

// CreatorEqualFold applies the EqualFold predicate on the "creator" field.
func CreatorEqualFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldEqualFold(FieldCreator, v))
}

// CreatorContainsFold applies the ContainsFold predicate on the "creator" field.
func CreatorContainsFold(v string) predicate.Turtle {
	return predicate.Turtle(sql.FieldContainsFold(FieldCreator, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v int32) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldState, v))
}

// CtimeEQ applies the EQ predicate on the "ctime" field.
func CtimeEQ(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldCtime, v))
}

// CtimeNEQ applies the NEQ predicate on the "ctime" field.
func CtimeNEQ(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldCtime, v))
}

// CtimeIn applies the In predicate on the "ctime" field.
func CtimeIn(vs ...time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldCtime, vs...))
}

// CtimeNotIn applies the NotIn predicate on the "ctime" field.
func CtimeNotIn(vs ...time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldCtime, vs...))
}

// CtimeGT applies the GT predicate on the "ctime" field.
func CtimeGT(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldCtime, v))
}

// CtimeGTE applies the GTE predicate on the "ctime" field.
func CtimeGTE(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldCtime, v))
}

// CtimeLT applies the LT predicate on the "ctime" field.
func CtimeLT(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldCtime, v))
}

// CtimeLTE applies the LTE predicate on the "ctime" field.
func CtimeLTE(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldCtime, v))
}

// MtimeEQ applies the EQ predicate on the "mtime" field.
func MtimeEQ(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldEQ(FieldMtime, v))
}

// MtimeNEQ applies the NEQ predicate on the "mtime" field.
func MtimeNEQ(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldNEQ(FieldMtime, v))
}

// MtimeIn applies the In predicate on the "mtime" field.
func MtimeIn(vs ...time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldIn(FieldMtime, vs...))
}

// MtimeNotIn applies the NotIn predicate on the "mtime" field.
func MtimeNotIn(vs ...time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldNotIn(FieldMtime, vs...))
}

// MtimeGT applies the GT predicate on the "mtime" field.
func MtimeGT(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldGT(FieldMtime, v))
}

// MtimeGTE applies the GTE predicate on the "mtime" field.
func MtimeGTE(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldGTE(FieldMtime, v))
}

// MtimeLT applies the LT predicate on the "mtime" field.
func MtimeLT(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldLT(FieldMtime, v))
}

// MtimeLTE applies the LTE predicate on the "mtime" field.
func MtimeLTE(v time.Time) predicate.Turtle {
	return predicate.Turtle(sql.FieldLTE(FieldMtime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Turtle) predicate.Turtle {
	return predicate.Turtle(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Turtle) predicate.Turtle {
	return predicate.Turtle(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Turtle) predicate.Turtle {
	return predicate.Turtle(sql.NotPredicates(p))
}
