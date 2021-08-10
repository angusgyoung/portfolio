import axios, { AxiosInstance } from 'axios';

const instance: AxiosInstance = axios.create()

const query = `
{
    user(login: "angusgyoung") {
        avatarUrl
        bio
        company
        isHireable
        
        organizations(first: 4) {
            nodes {
                avatarUrl
                login
                name
            }
        }
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
        publicKeys(first: 1) {
            nodes {
                createdAt
                fingerprint
                key
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
}`

export async function getUserDetails(): Promise<Record<string, any>> {
    try {
        const response = await instance({
            url: "https://api.github.com/graphql",
            method: "post",
            data: {
                query
            },
            headers: {
                "Authorization": `Bearer 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`,
                "Content-Type": "application/json"
            }
        })
        const data = await response.data

        return data.data.user
    } catch (error) {
        console.log(error)
        return Promise.reject(error)
    }
}