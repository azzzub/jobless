import jwt from 'jsonwebtoken'

/**
 * This function will check the token on local storage
 * whether the token is valid or not.
 * If the token was valid, than return. If not,
 * delete the token on local storage
 * @returns void
 */
const tokenChecker = (): void => {
  try {
    const tokenStored = localStorage.getItem('token')
    if (!tokenStored) return
    const validation = jwt.verify(tokenStored, process.env.NEXT_PUBLIC_JWT_SECRET)
    if (!validation) {
      localStorage.removeItem('token')
    }
  } catch (error) {
    localStorage.removeItem('token')
  }
}

export default tokenChecker
