import { useQuery, gql } from '@apollo/client'
import Link from 'next/link'

const GET_PROJECTS = gql`
  query {
    projects {
      slug
      creator {
        username
      }
      bids {
        ID
      }
      name
      price
      deadline
    }
  }
`

const App: React.FC = () => {
  const { loading, data, error } = useQuery<Projects>(GET_PROJECTS)

  if (loading) return <p>Loading...</p>
  if (error) return <p>Error</p>

  return (
    <div>
      <div>
        {data?.projects.map((element, i) => {
          return (
            <div key={i}>
              <div>Author: {element.creator.username}</div>
              <div>Judul: {element.name}</div>
              <div>Harga: {element.price}</div>
              <div>Deadline: {element.deadline}</div>
              <div>Penawar: {element.bids.length}</div>
              <Link href={`/p/${element.slug}`}>
                <button>Detail</button>
              </Link>
            </div>
          )
        })}
      </div>
    </div>
  )
}

export default App
