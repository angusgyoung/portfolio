package github

const userDataQuery = `
{
	user(login: "angusgyoung") {
		name
		login
		avatarUrl
		bio
		company
		isHireable
		url

		organizations(first: 6) {
			nodes {
				avatarUrl
				login
				name
				url
			}
		}
		
		publicKeys(first: 1) {
			nodes {
				createdAt
				fingerprint
				key
			}
		}
	  }
  }  
`

const projectDataQuery = `
{
	user(login: "angusgyoung") {
		repositoriesContributedTo(first: 6) {
			nodes {
				name
				description
				url
				updatedAt
				repositoryTopics (first: 10) {
					nodes {
						topic {
							name
						}
					}
				}
				stargazerCount
				releases (first: 4) {
					nodes {
					  name
					  tagName
					  updatedAt
					}
				}
				primaryLanguage {
					  name
					  color
				}
				languages (first: 4) {
					nodes {
						name
						color
					}
				}
			}
		}
		pinnedItems(first: 6, types: REPOSITORY) {
		  nodes {
			... on Repository {
			  name
			  description
			  url
			  updatedAt
			  repositoryTopics (first: 10) {
				  nodes {
					  topic {
						  name
					  }
				  }
			  }
			  stargazerCount
			  releases (first: 4) {
				  nodes {
					name
					tagName
					updatedAt
				  }
			  }
			  primaryLanguage {
					name
					color
			  }
			  languages (first: 4) {
				  nodes {
					  name
					  color
				  }
			  }
			}
		  }
	  }
	}
  }  
`