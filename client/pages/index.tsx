import { useQuery, gql } from '@apollo/client'
import Link from 'next/link'
import CreateProject from '../components/createProject'
import Login from '../components/login'
import Register from '../components/register'

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
  const { loading, data, error, refetch } = useQuery<Projects>(GET_PROJECTS)

  if (loading) return <div>Loading...</div>
  if (error) return <div>Error</div>

  return (
    <div>
      <Login />
      <Register />
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
        <CreateProject callback={() => refetch()} />
      </div>
    </div>
  )
}

export default App
