// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BugsColumns holds the columns for the "bugs" table.
	BugsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "jira_key", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "resolved_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "resolved", Type: field.TypeBool, Default: false},
		{Name: "priority", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "resolution_time", Type: field.TypeFloat64, Default: 0},
		{Name: "status", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "summary", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "project_key", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "assignment_time", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "prioritization_time", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "days_without_assignee", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "days_without_priority", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "days_without_resolution", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "days_without_component", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "labels", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "component", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "assignee", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "age", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "teams_bugs", Type: field.TypeUUID, Nullable: true},
	}
	// BugsTable holds the schema information for the "bugs" table.
	BugsTable = &schema.Table{
		Name:       "bugs",
		Columns:    BugsColumns,
		PrimaryKey: []*schema.Column{BugsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bugs_teams_bugs",
				Columns:    []*schema.Column{BugsColumns[22]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CodeCovsColumns holds the columns for the "code_covs" table.
	CodeCovsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "repository_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_organization", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "coverage_percentage", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "average_retests", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "average_retests_to_merge", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "coverage_trend", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_codecov", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// CodeCovsTable holds the schema information for the "code_covs" table.
	CodeCovsTable = &schema.Table{
		Name:       "code_covs",
		Columns:    CodeCovsColumns,
		PrimaryKey: []*schema.Column{CodeCovsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "code_covs_repositories_codecov",
				Columns:    []*schema.Column{CodeCovsColumns[7]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FailuresColumns holds the columns for the "failures" table.
	FailuresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "jira_key", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "jira_status", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "error_message", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "title_from_jira", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "closed_date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "labels", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "teams_failures", Type: field.TypeUUID, Nullable: true},
	}
	// FailuresTable holds the schema information for the "failures" table.
	FailuresTable = &schema.Table{
		Name:       "failures",
		Columns:    FailuresColumns,
		PrimaryKey: []*schema.Column{FailuresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "failures_teams_failures",
				Columns:    []*schema.Column{FailuresColumns[8]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProwJobsColumns holds the columns for the "prow_jobs" table.
	ProwJobsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "job_id", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "duration", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "tests_count", Type: field.TypeInt64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "failed_count", Type: field.TypeInt64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "skipped_count", Type: field.TypeInt64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_type", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "ci_failed", Type: field.TypeInt16, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "external_services_impact", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "e2e_failed_test_messages", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "suites_xml_url", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "build_error_logs", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_prow_jobs", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// ProwJobsTable holds the schema information for the "prow_jobs" table.
	ProwJobsTable = &schema.Table{
		Name:       "prow_jobs",
		Columns:    ProwJobsColumns,
		PrimaryKey: []*schema.Column{ProwJobsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prow_jobs_repositories_prow_jobs",
				Columns:    []*schema.Column{ProwJobsColumns[16]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProwSuitesColumns holds the columns for the "prow_suites" table.
	ProwSuitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "job_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_url", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "suite_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "status", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "error_message", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "external_services_impact", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "time", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "repository_prow_suites", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// ProwSuitesTable holds the schema information for the "prow_suites" table.
	ProwSuitesTable = &schema.Table{
		Name:       "prow_suites",
		Columns:    ProwSuitesColumns,
		PrimaryKey: []*schema.Column{ProwSuitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prow_suites_repositories_prow_suites",
				Columns:    []*schema.Column{ProwSuitesColumns[11]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PullRequestsColumns holds the columns for the "pull_requests" table.
	PullRequestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pr_id", Type: field.TypeUUID, Unique: true},
		{Name: "repository_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_organization", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "number", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "closed_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "merged_at", Type: field.TypeTime, SchemaType: map[string]string{"postgres": "timestamptz"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "author", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "title", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "merge_commit", Type: field.TypeString, Nullable: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "retest_count", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "retest_before_merge_count", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_prs", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// PullRequestsTable holds the schema information for the "pull_requests" table.
	PullRequestsTable = &schema.Table{
		Name:       "pull_requests",
		Columns:    PullRequestsColumns,
		PrimaryKey: []*schema.Column{PullRequestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pull_requests_repositories_prs",
				Columns:    []*schema.Column{PullRequestsColumns[14]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RepositoriesColumns holds the columns for the "repositories" table.
	RepositoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "repository_name", Type: field.TypeString, Unique: true, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_organization", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "description", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "git_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "teams_repositories", Type: field.TypeUUID, Nullable: true},
	}
	// RepositoriesTable holds the schema information for the "repositories" table.
	RepositoriesTable = &schema.Table{
		Name:       "repositories",
		Columns:    RepositoriesColumns,
		PrimaryKey: []*schema.Column{RepositoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "repositories_teams_repositories",
				Columns:    []*schema.Column{RepositoriesColumns[5]},
				RefColumns: []*schema.Column{TeamsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TeamsColumns holds the columns for the "teams" table.
	TeamsColumns = []*schema.Column{
		{Name: "team_id", Type: field.TypeUUID, Unique: true},
		{Name: "team_name", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString, Unique: true},
		{Name: "jira_keys", Type: field.TypeString},
	}
	// TeamsTable holds the schema information for the "teams" table.
	TeamsTable = &schema.Table{
		Name:       "teams",
		Columns:    TeamsColumns,
		PrimaryKey: []*schema.Column{TeamsColumns[0]},
	}
	// WorkflowsColumns holds the columns for the "workflows" table.
	WorkflowsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "workflow_id", Type: field.TypeUUID, Unique: true},
		{Name: "workflow_name", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "badge_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "html_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "job_url", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "state", Type: field.TypeString, Size: 2147483647, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "repository_workflows", Type: field.TypeString, Nullable: true, Size: 25},
	}
	// WorkflowsTable holds the schema information for the "workflows" table.
	WorkflowsTable = &schema.Table{
		Name:       "workflows",
		Columns:    WorkflowsColumns,
		PrimaryKey: []*schema.Column{WorkflowsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "workflows_repositories_workflows",
				Columns:    []*schema.Column{WorkflowsColumns[7]},
				RefColumns: []*schema.Column{RepositoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BugsTable,
		CodeCovsTable,
		FailuresTable,
		ProwJobsTable,
		ProwSuitesTable,
		PullRequestsTable,
		RepositoriesTable,
		TeamsTable,
		WorkflowsTable,
	}
)

func init() {
	BugsTable.ForeignKeys[0].RefTable = TeamsTable
	CodeCovsTable.ForeignKeys[0].RefTable = RepositoriesTable
	FailuresTable.ForeignKeys[0].RefTable = TeamsTable
	ProwJobsTable.ForeignKeys[0].RefTable = RepositoriesTable
	ProwSuitesTable.ForeignKeys[0].RefTable = RepositoriesTable
	PullRequestsTable.ForeignKeys[0].RefTable = RepositoriesTable
	RepositoriesTable.ForeignKeys[0].RefTable = TeamsTable
	WorkflowsTable.ForeignKeys[0].RefTable = RepositoriesTable
}
