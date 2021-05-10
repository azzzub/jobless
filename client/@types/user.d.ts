interface User {
  ID: number
  username: string
  email: string
  is_email_verified: boolean
  is_user_verified: boolean
  password: string
  first_name: string
  last_name: string
  created_at: string
  updated_at: string
  deleted_at: string
}

interface Users {
  users: User[]
}
