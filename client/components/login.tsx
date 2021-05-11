import { gql, useMutation } from '@apollo/client'
import { FormEvent, useEffect, useState } from 'react'

const LOGIN_MUTATION = gql`
  mutation Login($uoe: String!, $password: String!) {
    login(input: { uoe: $uoe, password: $password }) {
      token
    }
  }
`

const Login: React.FC = () => {
  const [uoe, setUoe] = useState('')
  const [password, setPassword] = useState('')
  const [token, setToken] = useState<string>()
  const [loginMutation, { loading, error }] = useMutation<LoginResponse>(LOGIN_MUTATION)

  const loginHandler = async (e: FormEvent): Promise<void> => {
    e.preventDefault()
    try {
      const { data } = await loginMutation({
        variables: {
          uoe,
          password,
        },
      })
      setToken(data?.login.token)
    } catch (error) {
      // Clearing token information on local storage if the login is failed
      localStorage.removeItem('token')
    }
  }

  useEffect(() => {
    if (token !== undefined) {
      localStorage.setItem('token', token)
    }
  }, [token])

  return (
    <div className="login">
      <form onSubmit={loginHandler}>
        <input
          type="username"
          placeholder="username/email"
          onChange={(e) => setUoe(e.target.value)}
        />
        <input
          type="password"
          placeholder="password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit" disabled={loading}>
          Login
        </button>
        {error ? <div>{error.message}</div> : null}
      </form>
    </div>
  )
}

export default Login
