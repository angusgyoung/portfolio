import axios, { AxiosInstance } from 'axios'
import { User } from '@/types'

const instance: AxiosInstance = axios.create()

const userQuery = `
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

const projectQuery = `
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

export async function getUser(): Promise<User> {
  return request(userQuery)
}

export async function getProjects(): Promise<User> {
  return request(projectQuery)
}

async function request(query: string): Promise<User> {
  try {
    const response = await instance({
      url: 'https://api.github.com/graphql',
      method: 'post',
      data: {
        query
      },
      headers: {
        Authorization: 'Bearer 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',
        'Content-Type': 'application/json'
      }
    })
    const data = await response.data

    console.log(data.data.user)
    return data.data.user
  } catch (error) {
    console.log(error)
    return Promise.reject(error)
  }
}
