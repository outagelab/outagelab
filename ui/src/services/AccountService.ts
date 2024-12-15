import type { Account } from '@/models/Account'

export class ApiAuthErrror extends Error {}

export default class AccountService {
  async getAccount(): Promise<Account> {
    const response = await fetch(`/api/account`, {
      headers: {
        Authorization: `Bearer ${localStorage.token}`
      }
    })

    if (!response.ok) {
      // TODO: handle other error types
      throw new ApiAuthErrror()
    }

    const account = (await response.json()) as Account

    return account
  }

  async updateAccount(account: Account) {
    const response = await fetch(`/api/account`, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${localStorage.token}`
      },
      body: JSON.stringify(account)
    })

    if (response.status !== 200) {
      throw new Error('request failed with status ' + response.status)
    }
  }

  async login(googleToken: string): Promise<string> {
    const loginResult = await fetch(`/api/login`, {
      method: 'POST',
      body: JSON.stringify({
        authProvider: 'google',
        authToken: googleToken
      })
    }).then((x) => x.json())

    return loginResult.authToken as string
  }
}
