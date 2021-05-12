import { useRouter } from 'next/router'
import { useEffect } from 'react'

const AuthCallback: React.FC = () => {
  const router = useRouter()
  const { token } = router.query

  useEffect(() => {
    if (Array.isArray(token)) {
      localStorage.setItem('token', token[0])
    } else {
      if (token !== undefined) {
        localStorage.setItem('token', token)
      }
    }
    router.push('/')
  }, [token, router])

  return <div>Callback</div>
}

export default AuthCallback
