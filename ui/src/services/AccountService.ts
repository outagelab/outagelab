import type { Account } from '@/models/Account'

export default class AccountService {
  async getAccount(): Promise<Account> {
    const account = await fetch(`/api/account`, {
      headers: {
        Authorization: `Bearer ${localStorage.token}`
      }
    }).then((x) => x.json())

    return account as Account
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
