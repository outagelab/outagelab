export interface Account {
  apiKeys: AccountApiKey[]
  applications: Application[]
}

export interface Application {
  id: string
  environments: Environment[]
  rules: Rule[]
}

export interface Environment {
  id: string
  enabled: boolean
}

export interface Rule {
  id: string
  enabled: boolean
  type: 'send-http'
  host: string
  status: number | null
  duration: number | null
}

export interface AccountApiKey {
  key: string
}
