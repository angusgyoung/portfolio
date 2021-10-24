package github

import "github.com/hasura/go-graphql-client"

type (
	Repository struct {
		Name graphql.String `json:"name"`
		Description graphql.String `json:"description"`
		Url graphql.String `json:"url"`
		UpdatedAt graphql.String `json:"updatedAt"`
		StargazerCount graphql.Int `json:"stargazerCount"`
	
		RepositoryTopics struct {
			Nodes []struct {
				Topic struct {
					Name graphql.String `json:"name"`
				} `json:"topic"`
			} `json:"nodes"`
		} `json:"repositoryTopics" graphql:"repositoryTopics(first: 10)"`
	
		Releases struct {
			Nodes []struct {
				Name graphql.String `json:"name"`
				TagName graphql.String `json:"tagName"`
				UpdatedAt graphql.String `json:"updatedAt"`
			} `json:"nodes"`
		} `json:"releases" graphql:"releases(first: 3)"`
	
		PrimaryLanguage struct {
			Name graphql.String `json:"name"`
			Color graphql.String `json:"color"`
		} `json:"primaryLanguage"`
		
		Languages struct {
			Nodes []struct {
				Name graphql.String `json:"name"`
				Color graphql.String `json:"color"`
			} `json:"nodes"`
		} `json:"languages" graphql:"languages(first: 6, orderBy: { field: SIZE, direction: DESC})"`
	}
)

var userDataQuery struct {
	User struct {
		Name graphql.String `json:"name"`
		Login graphql.String `json:"login"`
		AvatarUrl graphql.String `json:"avatarUrl"`
		Bio graphql.String `json:"bio"`
		Company graphql.String `json:"company"`
		IsHireable graphql.Boolean `json:"isHireable"`
		Url graphql.String `json:"url"`

		Organizations struct {
			Nodes []struct {
				Name graphql.String `json:"name"`
				Login graphql.String `json:"login"`
				AvatarUrl graphql.String `json:"avatarUrl"`
				Url graphql.String `json:"url"`
			} `json:"nodes"`
		} `json:"organizations" graphql:"organizations(first: 6)"`

		PublicKeys struct {
			Nodes []struct {
				CreatedAt graphql.String `json:"createdAt"`
				Fingerprint graphql.String `json:"fingerPrint"`
				Key graphql.String `json:"key"`
			} `json:"nodes"`
		} `json:"publicKeys" graphql:"publicKeys(first: 1)"`
	} `json:"user" graphql:"user(login: \"angusgyoung\")"`
}

var projectDataQuery struct {
	User struct {
		RepositoriesContributedTo struct {
			Nodes []struct {
				Repository `json:"repository" graphql:"... on Repository"`
			} `json:"nodes"`
		} `json:"repositoriesContributedTo" graphql:"repositoriesContributedTo(first: 10, includeUserRepositories: true, orderBy: { field: UPDATED_AT, direction: DESC})"`

		PinnedItems struct {
			Nodes []struct {
				Repository `json:"repository" graphql:"... on Repository"`
			} `json:"nodes"`
		} `json:"pinnedItems" graphql:"pinnedItems(first: 6, types: REPOSITORY)"`
	} `json:"user" graphql:"user(login: \"angusgyoung\")"`
}
