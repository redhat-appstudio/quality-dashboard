// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/predicate"
	"github.com/konflux-ci/quality-dashboard/pkg/storage/ent/db/pullrequests"
)

// PullRequestsDelete is the builder for deleting a PullRequests entity.
type PullRequestsDelete struct {
	config
	hooks    []Hook
	mutation *PullRequestsMutation
}

// Where appends a list predicates to the PullRequestsDelete builder.
func (prd *PullRequestsDelete) Where(ps ...predicate.PullRequests) *PullRequestsDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *PullRequestsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, PullRequestsMutation](ctx, prd.sqlExec, prd.mutation, prd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *PullRequestsDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *PullRequestsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: pullrequests.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pullrequests.FieldID,
			},
		},
	}
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prd.mutation.done = true
	return affected, err
}

// PullRequestsDeleteOne is the builder for deleting a single PullRequests entity.
type PullRequestsDeleteOne struct {
	prd *PullRequestsDelete
}

// Where appends a list predicates to the PullRequestsDelete builder.
func (prdo *PullRequestsDeleteOne) Where(ps ...predicate.PullRequests) *PullRequestsDeleteOne {
	prdo.prd.mutation.Where(ps...)
	return prdo
}

// Exec executes the deletion query.
func (prdo *PullRequestsDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pullrequests.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *PullRequestsDeleteOne) ExecX(ctx context.Context) {
	if err := prdo.Exec(ctx); err != nil {
		panic(err)
	}
}
