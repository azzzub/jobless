import React, { useEffect } from 'react'
import { ApolloProvider } from '@apollo/client/react'
import { AppProps } from 'next/app'
import client from '../config/apolloConfig'
import tokenChecker from '../utils/tokenChecker'

const App: React.FC<AppProps> = ({ Component, pageProps }) => {
  useEffect(() => {
    tokenChecker()
  }, [])

  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  )
}

export default App
