import { gql, useMutation } from '@apollo/client'
import { useRouter } from 'next/router'
import { useEffect } from 'react'

const VERIFICATION_MUTATION = gql`
  mutation VerificationToken($token: String!) {
    tokenVerification(input: { token: $token }) {
      token
      refresh_token
    }
  }
`

const AuthCallback: React.FC = () => {
  const router = useRouter()
  const { token } = router.query
  const [verification, { loading, error }] =
    useMutation<TokenVerificationResponse>(VERIFICATION_MUTATION)

  useEffect(() => {
    const verificationHandler = async (): Promise<void> => {
      try {
        const { data } = await verification({
          variables: {
            token,
          },
        })
        if (data) {
          localStorage.setItem('token', data.tokenVerification.token)
          localStorage.setItem('refresh_token', data.tokenVerification.refresh_token)
        }
        router.push('/')
      } catch (error) {
        // TODO
      }
    }
    verificationHandler()
  }, [token, router, verification])

  if (loading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>

  return <div>Redirect...</div>
}

export default AuthCallback
