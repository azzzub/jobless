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

interface TokenVerificationResponse {
  tokenVerification: {
    token: string
    refresh_token: string
  }
}
