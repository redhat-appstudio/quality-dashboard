// Code generated by ent, DO NOT EDIT.

package failure

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldID, id))
}

// JiraKey applies equality check predicate on the "jira_key" field. It's identical to JiraKeyEQ.
func JiraKey(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldJiraKey, v))
}

// JiraStatus applies equality check predicate on the "jira_status" field. It's identical to JiraStatusEQ.
func JiraStatus(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldJiraStatus, v))
}

// ErrorMessage applies equality check predicate on the "error_message" field. It's identical to ErrorMessageEQ.
func ErrorMessage(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldErrorMessage, v))
}

// TitleFromJira applies equality check predicate on the "title_from_jira" field. It's identical to TitleFromJiraEQ.
func TitleFromJira(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldTitleFromJira, v))
}

// CreatedDate applies equality check predicate on the "created_date" field. It's identical to CreatedDateEQ.
func CreatedDate(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldCreatedDate, v))
}

// ClosedDate applies equality check predicate on the "closed_date" field. It's identical to ClosedDateEQ.
func ClosedDate(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldClosedDate, v))
}

// Labels applies equality check predicate on the "labels" field. It's identical to LabelsEQ.
func Labels(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldLabels, v))
}

// JiraKeyEQ applies the EQ predicate on the "jira_key" field.
func JiraKeyEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldJiraKey, v))
}

// JiraKeyNEQ applies the NEQ predicate on the "jira_key" field.
func JiraKeyNEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldJiraKey, v))
}

// JiraKeyIn applies the In predicate on the "jira_key" field.
func JiraKeyIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldJiraKey, vs...))
}

// JiraKeyNotIn applies the NotIn predicate on the "jira_key" field.
func JiraKeyNotIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldJiraKey, vs...))
}

// JiraKeyGT applies the GT predicate on the "jira_key" field.
func JiraKeyGT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldJiraKey, v))
}

// JiraKeyGTE applies the GTE predicate on the "jira_key" field.
func JiraKeyGTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldJiraKey, v))
}

// JiraKeyLT applies the LT predicate on the "jira_key" field.
func JiraKeyLT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldJiraKey, v))
}

// JiraKeyLTE applies the LTE predicate on the "jira_key" field.
func JiraKeyLTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldJiraKey, v))
}

// JiraKeyContains applies the Contains predicate on the "jira_key" field.
func JiraKeyContains(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContains(FieldJiraKey, v))
}

// JiraKeyHasPrefix applies the HasPrefix predicate on the "jira_key" field.
func JiraKeyHasPrefix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasPrefix(FieldJiraKey, v))
}

// JiraKeyHasSuffix applies the HasSuffix predicate on the "jira_key" field.
func JiraKeyHasSuffix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasSuffix(FieldJiraKey, v))
}

// JiraKeyEqualFold applies the EqualFold predicate on the "jira_key" field.
func JiraKeyEqualFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEqualFold(FieldJiraKey, v))
}

// JiraKeyContainsFold applies the ContainsFold predicate on the "jira_key" field.
func JiraKeyContainsFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContainsFold(FieldJiraKey, v))
}

// JiraStatusEQ applies the EQ predicate on the "jira_status" field.
func JiraStatusEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldJiraStatus, v))
}

// JiraStatusNEQ applies the NEQ predicate on the "jira_status" field.
func JiraStatusNEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldJiraStatus, v))
}

// JiraStatusIn applies the In predicate on the "jira_status" field.
func JiraStatusIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldJiraStatus, vs...))
}

// JiraStatusNotIn applies the NotIn predicate on the "jira_status" field.
func JiraStatusNotIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldJiraStatus, vs...))
}

// JiraStatusGT applies the GT predicate on the "jira_status" field.
func JiraStatusGT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldJiraStatus, v))
}

// JiraStatusGTE applies the GTE predicate on the "jira_status" field.
func JiraStatusGTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldJiraStatus, v))
}

// JiraStatusLT applies the LT predicate on the "jira_status" field.
func JiraStatusLT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldJiraStatus, v))
}

// JiraStatusLTE applies the LTE predicate on the "jira_status" field.
func JiraStatusLTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldJiraStatus, v))
}

// JiraStatusContains applies the Contains predicate on the "jira_status" field.
func JiraStatusContains(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContains(FieldJiraStatus, v))
}

// JiraStatusHasPrefix applies the HasPrefix predicate on the "jira_status" field.
func JiraStatusHasPrefix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasPrefix(FieldJiraStatus, v))
}

// JiraStatusHasSuffix applies the HasSuffix predicate on the "jira_status" field.
func JiraStatusHasSuffix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasSuffix(FieldJiraStatus, v))
}

// JiraStatusEqualFold applies the EqualFold predicate on the "jira_status" field.
func JiraStatusEqualFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEqualFold(FieldJiraStatus, v))
}

// JiraStatusContainsFold applies the ContainsFold predicate on the "jira_status" field.
func JiraStatusContainsFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContainsFold(FieldJiraStatus, v))
}

// ErrorMessageEQ applies the EQ predicate on the "error_message" field.
func ErrorMessageEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldErrorMessage, v))
}

// ErrorMessageNEQ applies the NEQ predicate on the "error_message" field.
func ErrorMessageNEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldErrorMessage, v))
}

// ErrorMessageIn applies the In predicate on the "error_message" field.
func ErrorMessageIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldErrorMessage, vs...))
}

// ErrorMessageNotIn applies the NotIn predicate on the "error_message" field.
func ErrorMessageNotIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldErrorMessage, vs...))
}

// ErrorMessageGT applies the GT predicate on the "error_message" field.
func ErrorMessageGT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldErrorMessage, v))
}

// ErrorMessageGTE applies the GTE predicate on the "error_message" field.
func ErrorMessageGTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldErrorMessage, v))
}

// ErrorMessageLT applies the LT predicate on the "error_message" field.
func ErrorMessageLT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldErrorMessage, v))
}

// ErrorMessageLTE applies the LTE predicate on the "error_message" field.
func ErrorMessageLTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldErrorMessage, v))
}

// ErrorMessageContains applies the Contains predicate on the "error_message" field.
func ErrorMessageContains(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContains(FieldErrorMessage, v))
}

// ErrorMessageHasPrefix applies the HasPrefix predicate on the "error_message" field.
func ErrorMessageHasPrefix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasPrefix(FieldErrorMessage, v))
}

// ErrorMessageHasSuffix applies the HasSuffix predicate on the "error_message" field.
func ErrorMessageHasSuffix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasSuffix(FieldErrorMessage, v))
}

// ErrorMessageEqualFold applies the EqualFold predicate on the "error_message" field.
func ErrorMessageEqualFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEqualFold(FieldErrorMessage, v))
}

// ErrorMessageContainsFold applies the ContainsFold predicate on the "error_message" field.
func ErrorMessageContainsFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContainsFold(FieldErrorMessage, v))
}

// TitleFromJiraEQ applies the EQ predicate on the "title_from_jira" field.
func TitleFromJiraEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldTitleFromJira, v))
}

// TitleFromJiraNEQ applies the NEQ predicate on the "title_from_jira" field.
func TitleFromJiraNEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldTitleFromJira, v))
}

// TitleFromJiraIn applies the In predicate on the "title_from_jira" field.
func TitleFromJiraIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldTitleFromJira, vs...))
}

// TitleFromJiraNotIn applies the NotIn predicate on the "title_from_jira" field.
func TitleFromJiraNotIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldTitleFromJira, vs...))
}

// TitleFromJiraGT applies the GT predicate on the "title_from_jira" field.
func TitleFromJiraGT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldTitleFromJira, v))
}

// TitleFromJiraGTE applies the GTE predicate on the "title_from_jira" field.
func TitleFromJiraGTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldTitleFromJira, v))
}

// TitleFromJiraLT applies the LT predicate on the "title_from_jira" field.
func TitleFromJiraLT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldTitleFromJira, v))
}

// TitleFromJiraLTE applies the LTE predicate on the "title_from_jira" field.
func TitleFromJiraLTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldTitleFromJira, v))
}

// TitleFromJiraContains applies the Contains predicate on the "title_from_jira" field.
func TitleFromJiraContains(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContains(FieldTitleFromJira, v))
}

// TitleFromJiraHasPrefix applies the HasPrefix predicate on the "title_from_jira" field.
func TitleFromJiraHasPrefix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasPrefix(FieldTitleFromJira, v))
}

// TitleFromJiraHasSuffix applies the HasSuffix predicate on the "title_from_jira" field.
func TitleFromJiraHasSuffix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasSuffix(FieldTitleFromJira, v))
}

// TitleFromJiraIsNil applies the IsNil predicate on the "title_from_jira" field.
func TitleFromJiraIsNil() predicate.Failure {
	return predicate.Failure(sql.FieldIsNull(FieldTitleFromJira))
}

// TitleFromJiraNotNil applies the NotNil predicate on the "title_from_jira" field.
func TitleFromJiraNotNil() predicate.Failure {
	return predicate.Failure(sql.FieldNotNull(FieldTitleFromJira))
}

// TitleFromJiraEqualFold applies the EqualFold predicate on the "title_from_jira" field.
func TitleFromJiraEqualFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEqualFold(FieldTitleFromJira, v))
}

// TitleFromJiraContainsFold applies the ContainsFold predicate on the "title_from_jira" field.
func TitleFromJiraContainsFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContainsFold(FieldTitleFromJira, v))
}

// CreatedDateEQ applies the EQ predicate on the "created_date" field.
func CreatedDateEQ(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldCreatedDate, v))
}

// CreatedDateNEQ applies the NEQ predicate on the "created_date" field.
func CreatedDateNEQ(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldCreatedDate, v))
}

// CreatedDateIn applies the In predicate on the "created_date" field.
func CreatedDateIn(vs ...time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldCreatedDate, vs...))
}

// CreatedDateNotIn applies the NotIn predicate on the "created_date" field.
func CreatedDateNotIn(vs ...time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldCreatedDate, vs...))
}

// CreatedDateGT applies the GT predicate on the "created_date" field.
func CreatedDateGT(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldCreatedDate, v))
}

// CreatedDateGTE applies the GTE predicate on the "created_date" field.
func CreatedDateGTE(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldCreatedDate, v))
}

// CreatedDateLT applies the LT predicate on the "created_date" field.
func CreatedDateLT(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldCreatedDate, v))
}

// CreatedDateLTE applies the LTE predicate on the "created_date" field.
func CreatedDateLTE(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldCreatedDate, v))
}

// CreatedDateIsNil applies the IsNil predicate on the "created_date" field.
func CreatedDateIsNil() predicate.Failure {
	return predicate.Failure(sql.FieldIsNull(FieldCreatedDate))
}

// CreatedDateNotNil applies the NotNil predicate on the "created_date" field.
func CreatedDateNotNil() predicate.Failure {
	return predicate.Failure(sql.FieldNotNull(FieldCreatedDate))
}

// ClosedDateEQ applies the EQ predicate on the "closed_date" field.
func ClosedDateEQ(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldClosedDate, v))
}

// ClosedDateNEQ applies the NEQ predicate on the "closed_date" field.
func ClosedDateNEQ(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldClosedDate, v))
}

// ClosedDateIn applies the In predicate on the "closed_date" field.
func ClosedDateIn(vs ...time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldClosedDate, vs...))
}

// ClosedDateNotIn applies the NotIn predicate on the "closed_date" field.
func ClosedDateNotIn(vs ...time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldClosedDate, vs...))
}

// ClosedDateGT applies the GT predicate on the "closed_date" field.
func ClosedDateGT(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldClosedDate, v))
}

// ClosedDateGTE applies the GTE predicate on the "closed_date" field.
func ClosedDateGTE(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldClosedDate, v))
}

// ClosedDateLT applies the LT predicate on the "closed_date" field.
func ClosedDateLT(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldClosedDate, v))
}

// ClosedDateLTE applies the LTE predicate on the "closed_date" field.
func ClosedDateLTE(v time.Time) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldClosedDate, v))
}

// ClosedDateIsNil applies the IsNil predicate on the "closed_date" field.
func ClosedDateIsNil() predicate.Failure {
	return predicate.Failure(sql.FieldIsNull(FieldClosedDate))
}

// ClosedDateNotNil applies the NotNil predicate on the "closed_date" field.
func ClosedDateNotNil() predicate.Failure {
	return predicate.Failure(sql.FieldNotNull(FieldClosedDate))
}

// LabelsEQ applies the EQ predicate on the "labels" field.
func LabelsEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEQ(FieldLabels, v))
}

// LabelsNEQ applies the NEQ predicate on the "labels" field.
func LabelsNEQ(v string) predicate.Failure {
	return predicate.Failure(sql.FieldNEQ(FieldLabels, v))
}

// LabelsIn applies the In predicate on the "labels" field.
func LabelsIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldIn(FieldLabels, vs...))
}

// LabelsNotIn applies the NotIn predicate on the "labels" field.
func LabelsNotIn(vs ...string) predicate.Failure {
	return predicate.Failure(sql.FieldNotIn(FieldLabels, vs...))
}

// LabelsGT applies the GT predicate on the "labels" field.
func LabelsGT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGT(FieldLabels, v))
}

// LabelsGTE applies the GTE predicate on the "labels" field.
func LabelsGTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldGTE(FieldLabels, v))
}

// LabelsLT applies the LT predicate on the "labels" field.
func LabelsLT(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLT(FieldLabels, v))
}

// LabelsLTE applies the LTE predicate on the "labels" field.
func LabelsLTE(v string) predicate.Failure {
	return predicate.Failure(sql.FieldLTE(FieldLabels, v))
}

// LabelsContains applies the Contains predicate on the "labels" field.
func LabelsContains(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContains(FieldLabels, v))
}

// LabelsHasPrefix applies the HasPrefix predicate on the "labels" field.
func LabelsHasPrefix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasPrefix(FieldLabels, v))
}

// LabelsHasSuffix applies the HasSuffix predicate on the "labels" field.
func LabelsHasSuffix(v string) predicate.Failure {
	return predicate.Failure(sql.FieldHasSuffix(FieldLabels, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.Failure {
	return predicate.Failure(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.Failure {
	return predicate.Failure(sql.FieldNotNull(FieldLabels))
}

// LabelsEqualFold applies the EqualFold predicate on the "labels" field.
func LabelsEqualFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldEqualFold(FieldLabels, v))
}

// LabelsContainsFold applies the ContainsFold predicate on the "labels" field.
func LabelsContainsFold(v string) predicate.Failure {
	return predicate.Failure(sql.FieldContainsFold(FieldLabels, v))
}

// HasFailures applies the HasEdge predicate on the "failures" edge.
func HasFailures() predicate.Failure {
	return predicate.Failure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FailuresTable, FailuresColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFailuresWith applies the HasEdge predicate on the "failures" edge with a given conditions (other predicates).
func HasFailuresWith(preds ...predicate.Teams) predicate.Failure {
	return predicate.Failure(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FailuresInverseTable, TeamsFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, FailuresTable, FailuresColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Failure) predicate.Failure {
	return predicate.Failure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Failure) predicate.Failure {
	return predicate.Failure(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Failure) predicate.Failure {
	return predicate.Failure(func(s *sql.Selector) {
		p(s.Not())
	})
}
