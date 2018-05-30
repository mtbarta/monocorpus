interface userInfo {
  firstName: string
  lastName: string

  email: string
  password: string

  teams: string[]
  account: string

  isAdmin: boolean
}

export const emptyUserData = () => {
  return new UserData({
    firstName: '',
    lastName: '',
    email: '',
    password: '',
    teams: [] as string[],
    account: '',
    isAdmin: false
  })
}

export default class UserData {
  firstName: string
  lastName: string

  email: string
  password: string

  teams: string[]
  account: string

  isAdmin: boolean

  constructor(opts: userInfo) {
    // console.log(opts)
    this.firstName = opts.firstName
    this.lastName = opts.lastName

    this.email = opts.email
    this.password = opts.password

    this.teams = opts.teams
    this.account = opts.account

    this.isAdmin = opts.isAdmin
  }
}