import React from 'react'
import { ApolloProvider } from '@apollo/client/react'
import { AppProps } from 'next/app'
import client from '../config/apolloConfig'

const App: React.FC<AppProps> = ({ Component, pageProps }) => (
  <ApolloProvider client={client}>
    <Component {...pageProps} />
  </ApolloProvider>
)

export default App
