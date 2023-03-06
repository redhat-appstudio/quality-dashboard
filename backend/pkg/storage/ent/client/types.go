package client

import (
	coverageV1Alpha1 "github.com/redhat-appstudio/quality-studio/api/apis/codecov/v1alpha1"
	repoV1Alpha1 "github.com/redhat-appstudio/quality-studio/api/apis/github/v1alpha1"
	"github.com/redhat-appstudio/quality-studio/pkg/storage"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db"
)

func toStorageRepository(p *db.Repository) repoV1Alpha1.Repository {
	return repoV1Alpha1.Repository{
		Name: p.RepositoryName,
		Owner: repoV1Alpha1.Owner{
			Login: p.GitOrganization,
		},
		Description: p.Description,
		URL:         p.GitURL,
		ID:          p.ID,
	}
}

func toStorageWorkflows(p *db.Workflows) repoV1Alpha1.Workflow {
	return repoV1Alpha1.Workflow{
		Name:     p.WorkflowName,
		BadgeURL: p.BadgeURL,
		HTMLURL:  p.HTMLURL,
		State:    p.State,
	}
}

func toStorageRepositoryAllInfo(p *db.Repository, c *db.CodeCov) storage.RepositoryQualityInfo {
	return storage.RepositoryQualityInfo{
		GitOrganization: p.GitOrganization,
		RepositoryName:  p.RepositoryName,
		GitURL:          p.GitURL,
		Description:     p.Description,
		CodeCoverage: coverageV1Alpha1.Coverage{
			RepositoryName:     p.RepositoryName,
			GitOrganization:    p.GitOrganization,
			CoveragePercentage: c.CoveragePercentage,
		},
	}
}

func toStoragePrs(pr *db.PullRequests) repoV1Alpha1.PullRequest {
	return repoV1Alpha1.PullRequest{
		Title:     pr.Title,
		CreatedAt: pr.CreatedAt,
		MergedAt:  pr.MergedAt,
		ClosedAt:  pr.ClosedAt,
		State:     pr.State,
	}
}
