import { gql, useMutation } from '@apollo/client'
import { FormEvent, useEffect, useState } from 'react'

const REGISTER_MUTATION = gql`
  mutation Register(
    $first_name: String!
    $last_name: String!
    $username: String!
    $email: String!
    $password: String!
  ) {
    register(
      input: {
        first_name: $first_name
        last_name: $last_name
        username: $username
        email: $email
        password: $password
      }
    ) {
      ID
      first_name
      last_name
      username
      email
      created_at
    }
  }
`

const Register: React.FC = () => {
  const [firstName, setFirstName] = useState('')
  const [lastName, setLastName] = useState('')
  const [username, setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [response, setResponse] = useState<User>()
  const [registerMutation, { loading, error }] = useMutation<RegisterResponse>(REGISTER_MUTATION)

  const registerHandler = async (e: FormEvent): Promise<void> => {
    e.preventDefault()
    try {
      const { data } = await registerMutation({
        variables: {
          first_name: firstName,
          last_name: lastName,
          username,
          email,
          password,
        },
      })

      setResponse(data?.register)
    } catch (error) {
      // TODO
    }
  }

  useEffect(() => {
    // TODO
  }, [response])

  return (
    <div>
      <form onSubmit={registerHandler}>
        <input
          type="name"
          placeholder="nama depan"
          onChange={(e) => setFirstName(e.target.value)}
        />
        <input
          type="name"
          placeholder="nama belakang"
          onChange={(e) => setLastName(e.target.value)}
        />
        <input
          type="username"
          placeholder="username"
          onChange={(e) => setUsername(e.target.value)}
        />
        <input type="email" placeholder="email" onChange={(e) => setEmail(e.target.value)} />
        <input
          type="password"
          placeholder="password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <button type="submit" disabled={loading}>
          Register
        </button>
        {error ? <div>Error</div> : null}
      </form>
    </div>
  )
}

export default Register
