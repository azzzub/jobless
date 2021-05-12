interface Login {
  uoe: string
  password: string
}

interface LoginResponse {
  login: {
    token: string
    refresh_token: string
  }
}
