// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/failure"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/teams"
)

// FailureCreate is the builder for creating a Failure entity.
type FailureCreate struct {
	config
	mutation *FailureMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetJiraKey sets the "jira_key" field.
func (fc *FailureCreate) SetJiraKey(s string) *FailureCreate {
	fc.mutation.SetJiraKey(s)
	return fc
}

// SetJiraStatus sets the "jira_status" field.
func (fc *FailureCreate) SetJiraStatus(s string) *FailureCreate {
	fc.mutation.SetJiraStatus(s)
	return fc
}

// SetErrorMessage sets the "error_message" field.
func (fc *FailureCreate) SetErrorMessage(s string) *FailureCreate {
	fc.mutation.SetErrorMessage(s)
	return fc
}

// SetTitleFromJira sets the "title_from_jira" field.
func (fc *FailureCreate) SetTitleFromJira(s string) *FailureCreate {
	fc.mutation.SetTitleFromJira(s)
	return fc
}

// SetNillableTitleFromJira sets the "title_from_jira" field if the given value is not nil.
func (fc *FailureCreate) SetNillableTitleFromJira(s *string) *FailureCreate {
	if s != nil {
		fc.SetTitleFromJira(*s)
	}
	return fc
}

// SetCreatedDate sets the "created_date" field.
func (fc *FailureCreate) SetCreatedDate(t time.Time) *FailureCreate {
	fc.mutation.SetCreatedDate(t)
	return fc
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (fc *FailureCreate) SetNillableCreatedDate(t *time.Time) *FailureCreate {
	if t != nil {
		fc.SetCreatedDate(*t)
	}
	return fc
}

// SetClosedDate sets the "closed_date" field.
func (fc *FailureCreate) SetClosedDate(t time.Time) *FailureCreate {
	fc.mutation.SetClosedDate(t)
	return fc
}

// SetNillableClosedDate sets the "closed_date" field if the given value is not nil.
func (fc *FailureCreate) SetNillableClosedDate(t *time.Time) *FailureCreate {
	if t != nil {
		fc.SetClosedDate(*t)
	}
	return fc
}

// SetLabels sets the "labels" field.
func (fc *FailureCreate) SetLabels(s string) *FailureCreate {
	fc.mutation.SetLabels(s)
	return fc
}

// SetNillableLabels sets the "labels" field if the given value is not nil.
func (fc *FailureCreate) SetNillableLabels(s *string) *FailureCreate {
	if s != nil {
		fc.SetLabels(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FailureCreate) SetID(u uuid.UUID) *FailureCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FailureCreate) SetNillableID(u *uuid.UUID) *FailureCreate {
	if u != nil {
		fc.SetID(*u)
	}
	return fc
}

// SetFailuresID sets the "failures" edge to the Teams entity by ID.
func (fc *FailureCreate) SetFailuresID(id uuid.UUID) *FailureCreate {
	fc.mutation.SetFailuresID(id)
	return fc
}

// SetNillableFailuresID sets the "failures" edge to the Teams entity by ID if the given value is not nil.
func (fc *FailureCreate) SetNillableFailuresID(id *uuid.UUID) *FailureCreate {
	if id != nil {
		fc = fc.SetFailuresID(*id)
	}
	return fc
}

// SetFailures sets the "failures" edge to the Teams entity.
func (fc *FailureCreate) SetFailures(t *Teams) *FailureCreate {
	return fc.SetFailuresID(t.ID)
}

// Mutation returns the FailureMutation object of the builder.
func (fc *FailureCreate) Mutation() *FailureMutation {
	return fc.mutation
}

// Save creates the Failure in the database.
func (fc *FailureCreate) Save(ctx context.Context) (*Failure, error) {
	fc.defaults()
	return withHooks[*Failure, FailureMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FailureCreate) SaveX(ctx context.Context) *Failure {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FailureCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FailureCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FailureCreate) defaults() {
	if _, ok := fc.mutation.ID(); !ok {
		v := failure.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FailureCreate) check() error {
	if _, ok := fc.mutation.JiraKey(); !ok {
		return &ValidationError{Name: "jira_key", err: errors.New(`db: missing required field "Failure.jira_key"`)}
	}
	if v, ok := fc.mutation.JiraKey(); ok {
		if err := failure.JiraKeyValidator(v); err != nil {
			return &ValidationError{Name: "jira_key", err: fmt.Errorf(`db: validator failed for field "Failure.jira_key": %w`, err)}
		}
	}
	if _, ok := fc.mutation.JiraStatus(); !ok {
		return &ValidationError{Name: "jira_status", err: errors.New(`db: missing required field "Failure.jira_status"`)}
	}
	if _, ok := fc.mutation.ErrorMessage(); !ok {
		return &ValidationError{Name: "error_message", err: errors.New(`db: missing required field "Failure.error_message"`)}
	}
	return nil
}

func (fc *FailureCreate) sqlSave(ctx context.Context) (*Failure, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FailureCreate) createSpec() (*Failure, *sqlgraph.CreateSpec) {
	var (
		_node = &Failure{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: failure.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: failure.FieldID,
			},
		}
	)
	_spec.OnConflict = fc.conflict
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := fc.mutation.JiraKey(); ok {
		_spec.SetField(failure.FieldJiraKey, field.TypeString, value)
		_node.JiraKey = value
	}
	if value, ok := fc.mutation.JiraStatus(); ok {
		_spec.SetField(failure.FieldJiraStatus, field.TypeString, value)
		_node.JiraStatus = value
	}
	if value, ok := fc.mutation.ErrorMessage(); ok {
		_spec.SetField(failure.FieldErrorMessage, field.TypeString, value)
		_node.ErrorMessage = value
	}
	if value, ok := fc.mutation.TitleFromJira(); ok {
		_spec.SetField(failure.FieldTitleFromJira, field.TypeString, value)
		_node.TitleFromJira = &value
	}
	if value, ok := fc.mutation.CreatedDate(); ok {
		_spec.SetField(failure.FieldCreatedDate, field.TypeTime, value)
		_node.CreatedDate = &value
	}
	if value, ok := fc.mutation.ClosedDate(); ok {
		_spec.SetField(failure.FieldClosedDate, field.TypeTime, value)
		_node.ClosedDate = &value
	}
	if value, ok := fc.mutation.Labels(); ok {
		_spec.SetField(failure.FieldLabels, field.TypeString, value)
		_node.Labels = &value
	}
	if nodes := fc.mutation.FailuresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   failure.FailuresTable,
			Columns: []string{failure.FailuresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: teams.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.teams_failures = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Failure.Create().
//		SetJiraKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FailureUpsert) {
//			SetJiraKey(v+v).
//		}).
//		Exec(ctx)
func (fc *FailureCreate) OnConflict(opts ...sql.ConflictOption) *FailureUpsertOne {
	fc.conflict = opts
	return &FailureUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Failure.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FailureCreate) OnConflictColumns(columns ...string) *FailureUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FailureUpsertOne{
		create: fc,
	}
}

type (
	// FailureUpsertOne is the builder for "upsert"-ing
	//  one Failure node.
	FailureUpsertOne struct {
		create *FailureCreate
	}

	// FailureUpsert is the "OnConflict" setter.
	FailureUpsert struct {
		*sql.UpdateSet
	}
)

// SetJiraKey sets the "jira_key" field.
func (u *FailureUpsert) SetJiraKey(v string) *FailureUpsert {
	u.Set(failure.FieldJiraKey, v)
	return u
}

// UpdateJiraKey sets the "jira_key" field to the value that was provided on create.
func (u *FailureUpsert) UpdateJiraKey() *FailureUpsert {
	u.SetExcluded(failure.FieldJiraKey)
	return u
}

// SetJiraStatus sets the "jira_status" field.
func (u *FailureUpsert) SetJiraStatus(v string) *FailureUpsert {
	u.Set(failure.FieldJiraStatus, v)
	return u
}

// UpdateJiraStatus sets the "jira_status" field to the value that was provided on create.
func (u *FailureUpsert) UpdateJiraStatus() *FailureUpsert {
	u.SetExcluded(failure.FieldJiraStatus)
	return u
}

// SetErrorMessage sets the "error_message" field.
func (u *FailureUpsert) SetErrorMessage(v string) *FailureUpsert {
	u.Set(failure.FieldErrorMessage, v)
	return u
}

// UpdateErrorMessage sets the "error_message" field to the value that was provided on create.
func (u *FailureUpsert) UpdateErrorMessage() *FailureUpsert {
	u.SetExcluded(failure.FieldErrorMessage)
	return u
}

// SetTitleFromJira sets the "title_from_jira" field.
func (u *FailureUpsert) SetTitleFromJira(v string) *FailureUpsert {
	u.Set(failure.FieldTitleFromJira, v)
	return u
}

// UpdateTitleFromJira sets the "title_from_jira" field to the value that was provided on create.
func (u *FailureUpsert) UpdateTitleFromJira() *FailureUpsert {
	u.SetExcluded(failure.FieldTitleFromJira)
	return u
}

// ClearTitleFromJira clears the value of the "title_from_jira" field.
func (u *FailureUpsert) ClearTitleFromJira() *FailureUpsert {
	u.SetNull(failure.FieldTitleFromJira)
	return u
}

// SetCreatedDate sets the "created_date" field.
func (u *FailureUpsert) SetCreatedDate(v time.Time) *FailureUpsert {
	u.Set(failure.FieldCreatedDate, v)
	return u
}

// UpdateCreatedDate sets the "created_date" field to the value that was provided on create.
func (u *FailureUpsert) UpdateCreatedDate() *FailureUpsert {
	u.SetExcluded(failure.FieldCreatedDate)
	return u
}

// ClearCreatedDate clears the value of the "created_date" field.
func (u *FailureUpsert) ClearCreatedDate() *FailureUpsert {
	u.SetNull(failure.FieldCreatedDate)
	return u
}

// SetClosedDate sets the "closed_date" field.
func (u *FailureUpsert) SetClosedDate(v time.Time) *FailureUpsert {
	u.Set(failure.FieldClosedDate, v)
	return u
}

// UpdateClosedDate sets the "closed_date" field to the value that was provided on create.
func (u *FailureUpsert) UpdateClosedDate() *FailureUpsert {
	u.SetExcluded(failure.FieldClosedDate)
	return u
}

// ClearClosedDate clears the value of the "closed_date" field.
func (u *FailureUpsert) ClearClosedDate() *FailureUpsert {
	u.SetNull(failure.FieldClosedDate)
	return u
}

// SetLabels sets the "labels" field.
func (u *FailureUpsert) SetLabels(v string) *FailureUpsert {
	u.Set(failure.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *FailureUpsert) UpdateLabels() *FailureUpsert {
	u.SetExcluded(failure.FieldLabels)
	return u
}

// ClearLabels clears the value of the "labels" field.
func (u *FailureUpsert) ClearLabels() *FailureUpsert {
	u.SetNull(failure.FieldLabels)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Failure.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(failure.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FailureUpsertOne) UpdateNewValues() *FailureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(failure.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Failure.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FailureUpsertOne) Ignore() *FailureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FailureUpsertOne) DoNothing() *FailureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FailureCreate.OnConflict
// documentation for more info.
func (u *FailureUpsertOne) Update(set func(*FailureUpsert)) *FailureUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FailureUpsert{UpdateSet: update})
	}))
	return u
}

// SetJiraKey sets the "jira_key" field.
func (u *FailureUpsertOne) SetJiraKey(v string) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetJiraKey(v)
	})
}

// UpdateJiraKey sets the "jira_key" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateJiraKey() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateJiraKey()
	})
}

// SetJiraStatus sets the "jira_status" field.
func (u *FailureUpsertOne) SetJiraStatus(v string) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetJiraStatus(v)
	})
}

// UpdateJiraStatus sets the "jira_status" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateJiraStatus() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateJiraStatus()
	})
}

// SetErrorMessage sets the "error_message" field.
func (u *FailureUpsertOne) SetErrorMessage(v string) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetErrorMessage(v)
	})
}

// UpdateErrorMessage sets the "error_message" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateErrorMessage() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateErrorMessage()
	})
}

// SetTitleFromJira sets the "title_from_jira" field.
func (u *FailureUpsertOne) SetTitleFromJira(v string) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetTitleFromJira(v)
	})
}

// UpdateTitleFromJira sets the "title_from_jira" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateTitleFromJira() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateTitleFromJira()
	})
}

// ClearTitleFromJira clears the value of the "title_from_jira" field.
func (u *FailureUpsertOne) ClearTitleFromJira() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.ClearTitleFromJira()
	})
}

// SetCreatedDate sets the "created_date" field.
func (u *FailureUpsertOne) SetCreatedDate(v time.Time) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetCreatedDate(v)
	})
}

// UpdateCreatedDate sets the "created_date" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateCreatedDate() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateCreatedDate()
	})
}

// ClearCreatedDate clears the value of the "created_date" field.
func (u *FailureUpsertOne) ClearCreatedDate() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.ClearCreatedDate()
	})
}

// SetClosedDate sets the "closed_date" field.
func (u *FailureUpsertOne) SetClosedDate(v time.Time) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetClosedDate(v)
	})
}

// UpdateClosedDate sets the "closed_date" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateClosedDate() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateClosedDate()
	})
}

// ClearClosedDate clears the value of the "closed_date" field.
func (u *FailureUpsertOne) ClearClosedDate() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.ClearClosedDate()
	})
}

// SetLabels sets the "labels" field.
func (u *FailureUpsertOne) SetLabels(v string) *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *FailureUpsertOne) UpdateLabels() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *FailureUpsertOne) ClearLabels() *FailureUpsertOne {
	return u.Update(func(s *FailureUpsert) {
		s.ClearLabels()
	})
}

// Exec executes the query.
func (u *FailureUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for FailureCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FailureUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FailureUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: FailureUpsertOne.ID is not supported by MySQL driver. Use FailureUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FailureUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FailureCreateBulk is the builder for creating many Failure entities in bulk.
type FailureCreateBulk struct {
	config
	builders []*FailureCreate
	conflict []sql.ConflictOption
}

// Save creates the Failure entities in the database.
func (fcb *FailureCreateBulk) Save(ctx context.Context) ([]*Failure, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Failure, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FailureMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FailureCreateBulk) SaveX(ctx context.Context) []*Failure {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FailureCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FailureCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Failure.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FailureUpsert) {
//			SetJiraKey(v+v).
//		}).
//		Exec(ctx)
func (fcb *FailureCreateBulk) OnConflict(opts ...sql.ConflictOption) *FailureUpsertBulk {
	fcb.conflict = opts
	return &FailureUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Failure.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FailureCreateBulk) OnConflictColumns(columns ...string) *FailureUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FailureUpsertBulk{
		create: fcb,
	}
}

// FailureUpsertBulk is the builder for "upsert"-ing
// a bulk of Failure nodes.
type FailureUpsertBulk struct {
	create *FailureCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Failure.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(failure.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *FailureUpsertBulk) UpdateNewValues() *FailureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(failure.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Failure.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FailureUpsertBulk) Ignore() *FailureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FailureUpsertBulk) DoNothing() *FailureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FailureCreateBulk.OnConflict
// documentation for more info.
func (u *FailureUpsertBulk) Update(set func(*FailureUpsert)) *FailureUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FailureUpsert{UpdateSet: update})
	}))
	return u
}

// SetJiraKey sets the "jira_key" field.
func (u *FailureUpsertBulk) SetJiraKey(v string) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetJiraKey(v)
	})
}

// UpdateJiraKey sets the "jira_key" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateJiraKey() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateJiraKey()
	})
}

// SetJiraStatus sets the "jira_status" field.
func (u *FailureUpsertBulk) SetJiraStatus(v string) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetJiraStatus(v)
	})
}

// UpdateJiraStatus sets the "jira_status" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateJiraStatus() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateJiraStatus()
	})
}

// SetErrorMessage sets the "error_message" field.
func (u *FailureUpsertBulk) SetErrorMessage(v string) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetErrorMessage(v)
	})
}

// UpdateErrorMessage sets the "error_message" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateErrorMessage() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateErrorMessage()
	})
}

// SetTitleFromJira sets the "title_from_jira" field.
func (u *FailureUpsertBulk) SetTitleFromJira(v string) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetTitleFromJira(v)
	})
}

// UpdateTitleFromJira sets the "title_from_jira" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateTitleFromJira() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateTitleFromJira()
	})
}

// ClearTitleFromJira clears the value of the "title_from_jira" field.
func (u *FailureUpsertBulk) ClearTitleFromJira() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.ClearTitleFromJira()
	})
}

// SetCreatedDate sets the "created_date" field.
func (u *FailureUpsertBulk) SetCreatedDate(v time.Time) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetCreatedDate(v)
	})
}

// UpdateCreatedDate sets the "created_date" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateCreatedDate() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateCreatedDate()
	})
}

// ClearCreatedDate clears the value of the "created_date" field.
func (u *FailureUpsertBulk) ClearCreatedDate() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.ClearCreatedDate()
	})
}

// SetClosedDate sets the "closed_date" field.
func (u *FailureUpsertBulk) SetClosedDate(v time.Time) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetClosedDate(v)
	})
}

// UpdateClosedDate sets the "closed_date" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateClosedDate() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateClosedDate()
	})
}

// ClearClosedDate clears the value of the "closed_date" field.
func (u *FailureUpsertBulk) ClearClosedDate() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.ClearClosedDate()
	})
}

// SetLabels sets the "labels" field.
func (u *FailureUpsertBulk) SetLabels(v string) *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *FailureUpsertBulk) UpdateLabels() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *FailureUpsertBulk) ClearLabels() *FailureUpsertBulk {
	return u.Update(func(s *FailureUpsert) {
		s.ClearLabels()
	})
}

// Exec executes the query.
func (u *FailureUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the FailureCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for FailureCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FailureUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
