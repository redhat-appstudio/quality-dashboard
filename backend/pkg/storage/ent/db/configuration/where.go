// Code generated by ent, DO NOT EDIT.

package configuration

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldID, id))
}

// TeamName applies equality check predicate on the "team_name" field. It's identical to TeamNameEQ.
func TeamName(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldTeamName, v))
}

// JiraConfig applies equality check predicate on the "jira_config" field. It's identical to JiraConfigEQ.
func JiraConfig(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldJiraConfig, v))
}

// BugSlosConfig applies equality check predicate on the "bug_slos_config" field. It's identical to BugSlosConfigEQ.
func BugSlosConfig(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldBugSlosConfig, v))
}

// TeamNameEQ applies the EQ predicate on the "team_name" field.
func TeamNameEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldTeamName, v))
}

// TeamNameNEQ applies the NEQ predicate on the "team_name" field.
func TeamNameNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldTeamName, v))
}

// TeamNameIn applies the In predicate on the "team_name" field.
func TeamNameIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldTeamName, vs...))
}

// TeamNameNotIn applies the NotIn predicate on the "team_name" field.
func TeamNameNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldTeamName, vs...))
}

// TeamNameGT applies the GT predicate on the "team_name" field.
func TeamNameGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldTeamName, v))
}

// TeamNameGTE applies the GTE predicate on the "team_name" field.
func TeamNameGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldTeamName, v))
}

// TeamNameLT applies the LT predicate on the "team_name" field.
func TeamNameLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldTeamName, v))
}

// TeamNameLTE applies the LTE predicate on the "team_name" field.
func TeamNameLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldTeamName, v))
}

// TeamNameContains applies the Contains predicate on the "team_name" field.
func TeamNameContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldTeamName, v))
}

// TeamNameHasPrefix applies the HasPrefix predicate on the "team_name" field.
func TeamNameHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldTeamName, v))
}

// TeamNameHasSuffix applies the HasSuffix predicate on the "team_name" field.
func TeamNameHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldTeamName, v))
}

// TeamNameEqualFold applies the EqualFold predicate on the "team_name" field.
func TeamNameEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldTeamName, v))
}

// TeamNameContainsFold applies the ContainsFold predicate on the "team_name" field.
func TeamNameContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldTeamName, v))
}

// JiraConfigEQ applies the EQ predicate on the "jira_config" field.
func JiraConfigEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldJiraConfig, v))
}

// JiraConfigNEQ applies the NEQ predicate on the "jira_config" field.
func JiraConfigNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldJiraConfig, v))
}

// JiraConfigIn applies the In predicate on the "jira_config" field.
func JiraConfigIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldJiraConfig, vs...))
}

// JiraConfigNotIn applies the NotIn predicate on the "jira_config" field.
func JiraConfigNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldJiraConfig, vs...))
}

// JiraConfigGT applies the GT predicate on the "jira_config" field.
func JiraConfigGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldJiraConfig, v))
}

// JiraConfigGTE applies the GTE predicate on the "jira_config" field.
func JiraConfigGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldJiraConfig, v))
}

// JiraConfigLT applies the LT predicate on the "jira_config" field.
func JiraConfigLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldJiraConfig, v))
}

// JiraConfigLTE applies the LTE predicate on the "jira_config" field.
func JiraConfigLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldJiraConfig, v))
}

// JiraConfigContains applies the Contains predicate on the "jira_config" field.
func JiraConfigContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldJiraConfig, v))
}

// JiraConfigHasPrefix applies the HasPrefix predicate on the "jira_config" field.
func JiraConfigHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldJiraConfig, v))
}

// JiraConfigHasSuffix applies the HasSuffix predicate on the "jira_config" field.
func JiraConfigHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldJiraConfig, v))
}

// JiraConfigEqualFold applies the EqualFold predicate on the "jira_config" field.
func JiraConfigEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldJiraConfig, v))
}

// JiraConfigContainsFold applies the ContainsFold predicate on the "jira_config" field.
func JiraConfigContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldJiraConfig, v))
}

// BugSlosConfigEQ applies the EQ predicate on the "bug_slos_config" field.
func BugSlosConfigEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEQ(FieldBugSlosConfig, v))
}

// BugSlosConfigNEQ applies the NEQ predicate on the "bug_slos_config" field.
func BugSlosConfigNEQ(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNEQ(FieldBugSlosConfig, v))
}

// BugSlosConfigIn applies the In predicate on the "bug_slos_config" field.
func BugSlosConfigIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldIn(FieldBugSlosConfig, vs...))
}

// BugSlosConfigNotIn applies the NotIn predicate on the "bug_slos_config" field.
func BugSlosConfigNotIn(vs ...string) predicate.Configuration {
	return predicate.Configuration(sql.FieldNotIn(FieldBugSlosConfig, vs...))
}

// BugSlosConfigGT applies the GT predicate on the "bug_slos_config" field.
func BugSlosConfigGT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGT(FieldBugSlosConfig, v))
}

// BugSlosConfigGTE applies the GTE predicate on the "bug_slos_config" field.
func BugSlosConfigGTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldGTE(FieldBugSlosConfig, v))
}

// BugSlosConfigLT applies the LT predicate on the "bug_slos_config" field.
func BugSlosConfigLT(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLT(FieldBugSlosConfig, v))
}

// BugSlosConfigLTE applies the LTE predicate on the "bug_slos_config" field.
func BugSlosConfigLTE(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldLTE(FieldBugSlosConfig, v))
}

// BugSlosConfigContains applies the Contains predicate on the "bug_slos_config" field.
func BugSlosConfigContains(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContains(FieldBugSlosConfig, v))
}

// BugSlosConfigHasPrefix applies the HasPrefix predicate on the "bug_slos_config" field.
func BugSlosConfigHasPrefix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasPrefix(FieldBugSlosConfig, v))
}

// BugSlosConfigHasSuffix applies the HasSuffix predicate on the "bug_slos_config" field.
func BugSlosConfigHasSuffix(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldHasSuffix(FieldBugSlosConfig, v))
}

// BugSlosConfigEqualFold applies the EqualFold predicate on the "bug_slos_config" field.
func BugSlosConfigEqualFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldEqualFold(FieldBugSlosConfig, v))
}

// BugSlosConfigContainsFold applies the ContainsFold predicate on the "bug_slos_config" field.
func BugSlosConfigContainsFold(v string) predicate.Configuration {
	return predicate.Configuration(sql.FieldContainsFold(FieldBugSlosConfig, v))
}

// HasConfiguration applies the HasEdge predicate on the "configuration" edge.
func HasConfiguration() predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConfigurationTable, ConfigurationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigurationWith applies the HasEdge predicate on the "configuration" edge with a given conditions (other predicates).
func HasConfigurationWith(preds ...predicate.Teams) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConfigurationInverseTable, TeamsFieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConfigurationTable, ConfigurationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
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
func Not(p predicate.Configuration) predicate.Configuration {
	return predicate.Configuration(func(s *sql.Selector) {
		p(s.Not())
	})
}
