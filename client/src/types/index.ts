export interface Organisation {
  avatarUrl: string
  login: string
  name: string
  url: string
}

export interface Organizations {
  nodes: Organisation[]
}

export interface Language {
  name: string
  color: string
}

export interface Languages {
  nodes: Language[]
}

export interface PublicKey {
  createdAt: Date
  fingerprint: string
  key: string
}

export interface PublicKeys {
  nodes: PublicKey[]
}

export interface Release {
  name: string
  tagName: string
  updatedAt: Date
}

export interface Releases {
  nodes: Release[]
}

export interface Topic {
  name: string
}

export interface RepositoryTopics {
  nodes: Topic[]
}

export interface PrimaryLanguage {
  name: string
  color: string
}

export interface Repository {
  name: string
  description?: string
  url: string
  updatedAt: Date
  repositoryTopics: RepositoryTopics
  stargazerCount: number
  releases: Releases
  primaryLanguage: PrimaryLanguage
  languages: Languages
}

export interface PinnedItems {
  nodes: Repository[]
}

export interface Repositories {
  nodes: Repository[]
}

export interface User {
  name: string
  login: string
  avatarUrl: string
  bio: string
  company?: string
  isHireable: boolean
  url: string
  organizations: Organizations
  publicKeys: PublicKeys
  repositoriesContributedTo: Repositories
  pinnedItems: PinnedItems
}

export interface Data {
  user: User
}

export interface RootObject {
  data: Data
}
