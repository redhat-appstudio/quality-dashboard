// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/bugs"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/failure"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/predicate"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/repository"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/teams"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/configuration"
)

// TeamsUpdate is the builder for updating Teams entities.
type TeamsUpdate struct {
	config
	hooks    []Hook
	mutation *TeamsMutation
}

// Where appends a list predicates to the TeamsUpdate builder.
func (tu *TeamsUpdate) Where(ps ...predicate.Teams) *TeamsUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetTeamName sets the "team_name" field.
func (tu *TeamsUpdate) SetTeamName(s string) *TeamsUpdate {
	tu.mutation.SetTeamName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TeamsUpdate) SetDescription(s string) *TeamsUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetJiraKeys sets the "jira_keys" field.
func (tu *TeamsUpdate) SetJiraKeys(s string) *TeamsUpdate {
	tu.mutation.SetJiraKeys(s)
	return tu
}

// AddRepositoryIDs adds the "repositories" edge to the Repository entity by IDs.
func (tu *TeamsUpdate) AddRepositoryIDs(ids ...string) *TeamsUpdate {
	tu.mutation.AddRepositoryIDs(ids...)
	return tu
}

// AddRepositories adds the "repositories" edges to the Repository entity.
func (tu *TeamsUpdate) AddRepositories(r ...*Repository) *TeamsUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.AddRepositoryIDs(ids...)
}

// AddBugIDs adds the "bugs" edge to the Bugs entity by IDs.
func (tu *TeamsUpdate) AddBugIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.AddBugIDs(ids...)
	return tu
}

// AddBugs adds the "bugs" edges to the Bugs entity.
func (tu *TeamsUpdate) AddBugs(b ...*Bugs) *TeamsUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.AddBugIDs(ids...)
}

// AddFailureIDs adds the "failures" edge to the Failure entity by IDs.
func (tu *TeamsUpdate) AddFailureIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.AddFailureIDs(ids...)
	return tu
}

// AddFailures adds the "failures" edges to the Failure entity.
func (tu *TeamsUpdate) AddFailures(f ...*Failure) *TeamsUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.AddFailureIDs(ids...)
}

// AddConfigurationIDs adds the "configuration" edge to the Configuration entity by IDs.
func (tu *TeamsUpdate) AddConfigurationIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.AddConfigurationIDs(ids...)
	return tu
}

// AddConfiguration adds the "configuration" edges to the Configuration entity.
func (tu *TeamsUpdate) AddConfiguration(c ...*Configuration) *TeamsUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddConfigurationIDs(ids...)
}

// Mutation returns the TeamsMutation object of the builder.
func (tu *TeamsUpdate) Mutation() *TeamsMutation {
	return tu.mutation
}

// ClearRepositories clears all "repositories" edges to the Repository entity.
func (tu *TeamsUpdate) ClearRepositories() *TeamsUpdate {
	tu.mutation.ClearRepositories()
	return tu
}

// RemoveRepositoryIDs removes the "repositories" edge to Repository entities by IDs.
func (tu *TeamsUpdate) RemoveRepositoryIDs(ids ...string) *TeamsUpdate {
	tu.mutation.RemoveRepositoryIDs(ids...)
	return tu
}

// RemoveRepositories removes "repositories" edges to Repository entities.
func (tu *TeamsUpdate) RemoveRepositories(r ...*Repository) *TeamsUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tu.RemoveRepositoryIDs(ids...)
}

// ClearBugs clears all "bugs" edges to the Bugs entity.
func (tu *TeamsUpdate) ClearBugs() *TeamsUpdate {
	tu.mutation.ClearBugs()
	return tu
}

// RemoveBugIDs removes the "bugs" edge to Bugs entities by IDs.
func (tu *TeamsUpdate) RemoveBugIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.RemoveBugIDs(ids...)
	return tu
}

// RemoveBugs removes "bugs" edges to Bugs entities.
func (tu *TeamsUpdate) RemoveBugs(b ...*Bugs) *TeamsUpdate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.RemoveBugIDs(ids...)
}

// ClearFailures clears all "failures" edges to the Failure entity.
func (tu *TeamsUpdate) ClearFailures() *TeamsUpdate {
	tu.mutation.ClearFailures()
	return tu
}

// RemoveFailureIDs removes the "failures" edge to Failure entities by IDs.
func (tu *TeamsUpdate) RemoveFailureIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.RemoveFailureIDs(ids...)
	return tu
}

// RemoveFailures removes "failures" edges to Failure entities.
func (tu *TeamsUpdate) RemoveFailures(f ...*Failure) *TeamsUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tu.RemoveFailureIDs(ids...)
}

// ClearConfiguration clears all "configuration" edges to the Configuration entity.
func (tu *TeamsUpdate) ClearConfiguration() *TeamsUpdate {
	tu.mutation.ClearConfiguration()
	return tu
}

// RemoveConfigurationIDs removes the "configuration" edge to Configuration entities by IDs.
func (tu *TeamsUpdate) RemoveConfigurationIDs(ids ...uuid.UUID) *TeamsUpdate {
	tu.mutation.RemoveConfigurationIDs(ids...)
	return tu
}

// RemoveConfiguration removes "configuration" edges to Configuration entities.
func (tu *TeamsUpdate) RemoveConfiguration(c ...*Configuration) *TeamsUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveConfigurationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TeamsMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamsUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamsUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamsUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TeamsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teams.Table,
			Columns: teams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: teams.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.TeamName(); ok {
		_spec.SetField(teams.FieldTeamName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(teams.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.JiraKeys(); ok {
		_spec.SetField(teams.FieldJiraKeys, field.TypeString, value)
	}
	if tu.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedRepositoriesIDs(); len(nodes) > 0 && !tu.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.BugsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedBugsIDs(); len(nodes) > 0 && !tu.mutation.BugsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BugsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.FailuresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedFailuresIDs(); len(nodes) > 0 && !tu.mutation.FailuresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.FailuresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedConfigurationIDs(); len(nodes) > 0 && !tu.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ConfigurationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teams.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamsUpdateOne is the builder for updating a single Teams entity.
type TeamsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeamsMutation
}

// SetTeamName sets the "team_name" field.
func (tuo *TeamsUpdateOne) SetTeamName(s string) *TeamsUpdateOne {
	tuo.mutation.SetTeamName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TeamsUpdateOne) SetDescription(s string) *TeamsUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetJiraKeys sets the "jira_keys" field.
func (tuo *TeamsUpdateOne) SetJiraKeys(s string) *TeamsUpdateOne {
	tuo.mutation.SetJiraKeys(s)
	return tuo
}

// AddRepositoryIDs adds the "repositories" edge to the Repository entity by IDs.
func (tuo *TeamsUpdateOne) AddRepositoryIDs(ids ...string) *TeamsUpdateOne {
	tuo.mutation.AddRepositoryIDs(ids...)
	return tuo
}

// AddRepositories adds the "repositories" edges to the Repository entity.
func (tuo *TeamsUpdateOne) AddRepositories(r ...*Repository) *TeamsUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.AddRepositoryIDs(ids...)
}

// AddBugIDs adds the "bugs" edge to the Bugs entity by IDs.
func (tuo *TeamsUpdateOne) AddBugIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.AddBugIDs(ids...)
	return tuo
}

// AddBugs adds the "bugs" edges to the Bugs entity.
func (tuo *TeamsUpdateOne) AddBugs(b ...*Bugs) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.AddBugIDs(ids...)
}

// AddFailureIDs adds the "failures" edge to the Failure entity by IDs.
func (tuo *TeamsUpdateOne) AddFailureIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.AddFailureIDs(ids...)
	return tuo
}

// AddFailures adds the "failures" edges to the Failure entity.
func (tuo *TeamsUpdateOne) AddFailures(f ...*Failure) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.AddFailureIDs(ids...)
}

// AddConfigurationIDs adds the "configuration" edge to the Configuration entity by IDs.
func (tuo *TeamsUpdateOne) AddConfigurationIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.AddConfigurationIDs(ids...)
	return tuo
}

// AddConfiguration adds the "configuration" edges to the Configuration entity.
func (tuo *TeamsUpdateOne) AddConfiguration(c ...*Configuration) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddConfigurationIDs(ids...)
}

// Mutation returns the TeamsMutation object of the builder.
func (tuo *TeamsUpdateOne) Mutation() *TeamsMutation {
	return tuo.mutation
}

// ClearRepositories clears all "repositories" edges to the Repository entity.
func (tuo *TeamsUpdateOne) ClearRepositories() *TeamsUpdateOne {
	tuo.mutation.ClearRepositories()
	return tuo
}

// RemoveRepositoryIDs removes the "repositories" edge to Repository entities by IDs.
func (tuo *TeamsUpdateOne) RemoveRepositoryIDs(ids ...string) *TeamsUpdateOne {
	tuo.mutation.RemoveRepositoryIDs(ids...)
	return tuo
}

// RemoveRepositories removes "repositories" edges to Repository entities.
func (tuo *TeamsUpdateOne) RemoveRepositories(r ...*Repository) *TeamsUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return tuo.RemoveRepositoryIDs(ids...)
}

// ClearBugs clears all "bugs" edges to the Bugs entity.
func (tuo *TeamsUpdateOne) ClearBugs() *TeamsUpdateOne {
	tuo.mutation.ClearBugs()
	return tuo
}

// RemoveBugIDs removes the "bugs" edge to Bugs entities by IDs.
func (tuo *TeamsUpdateOne) RemoveBugIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.RemoveBugIDs(ids...)
	return tuo
}

// RemoveBugs removes "bugs" edges to Bugs entities.
func (tuo *TeamsUpdateOne) RemoveBugs(b ...*Bugs) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.RemoveBugIDs(ids...)
}

// ClearFailures clears all "failures" edges to the Failure entity.
func (tuo *TeamsUpdateOne) ClearFailures() *TeamsUpdateOne {
	tuo.mutation.ClearFailures()
	return tuo
}

// RemoveFailureIDs removes the "failures" edge to Failure entities by IDs.
func (tuo *TeamsUpdateOne) RemoveFailureIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.RemoveFailureIDs(ids...)
	return tuo
}

// RemoveFailures removes "failures" edges to Failure entities.
func (tuo *TeamsUpdateOne) RemoveFailures(f ...*Failure) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tuo.RemoveFailureIDs(ids...)
}

// ClearConfiguration clears all "configuration" edges to the Configuration entity.
func (tuo *TeamsUpdateOne) ClearConfiguration() *TeamsUpdateOne {
	tuo.mutation.ClearConfiguration()
	return tuo
}

// RemoveConfigurationIDs removes the "configuration" edge to Configuration entities by IDs.
func (tuo *TeamsUpdateOne) RemoveConfigurationIDs(ids ...uuid.UUID) *TeamsUpdateOne {
	tuo.mutation.RemoveConfigurationIDs(ids...)
	return tuo
}

// RemoveConfiguration removes "configuration" edges to Configuration entities.
func (tuo *TeamsUpdateOne) RemoveConfiguration(c ...*Configuration) *TeamsUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveConfigurationIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamsUpdateOne) Select(field string, fields ...string) *TeamsUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teams entity.
func (tuo *TeamsUpdateOne) Save(ctx context.Context) (*Teams, error) {
	return withHooks[*Teams, TeamsMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamsUpdateOne) SaveX(ctx context.Context) *Teams {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamsUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamsUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TeamsUpdateOne) sqlSave(ctx context.Context) (_node *Teams, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   teams.Table,
			Columns: teams.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: teams.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Teams.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teams.FieldID)
		for _, f := range fields {
			if !teams.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != teams.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TeamName(); ok {
		_spec.SetField(teams.FieldTeamName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(teams.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.JiraKeys(); ok {
		_spec.SetField(teams.FieldJiraKeys, field.TypeString, value)
	}
	if tuo.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedRepositoriesIDs(); len(nodes) > 0 && !tuo.mutation.RepositoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RepositoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.RepositoriesTable,
			Columns: []string{teams.RepositoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: repository.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.BugsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedBugsIDs(); len(nodes) > 0 && !tuo.mutation.BugsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BugsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.BugsTable,
			Columns: []string{teams.BugsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: bugs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.FailuresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedFailuresIDs(); len(nodes) > 0 && !tuo.mutation.FailuresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.FailuresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.FailuresTable,
			Columns: []string{teams.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: failure.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedConfigurationIDs(); len(nodes) > 0 && !tuo.mutation.ConfigurationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ConfigurationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teams.ConfigurationTable,
			Columns: []string{teams.ConfigurationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: configuration.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Teams{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teams.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
