import { useQuery, gql } from '@apollo/client'
import React from 'react'

const BIDS = gql`
  query {
    projects {
      ID
      price
    }
  }
`

const App: React.FC = () => {
  const { loading, data, error } = useQuery(BIDS)
  console.log(data)

  if (loading) return <p>Loading...</p>
  if (error) return <p>Error :(</p>

  return <div>test</div>
}

export default App
