import { gql, useQuery } from '@apollo/client'
import { useRouter } from 'next/router'

const GET_PROJECT_ON_SLUG = gql`
  query Project($slug: String!) {
    project(slug: $slug) {
      ID
      name
      desc
      price
      deadline
      creator {
        username
        first_name
        last_name
      }
      created_at
      updated_at
    }
  }
`

const ProjectDetails: React.FC = () => {
  const router = useRouter()
  const { slug } = router.query
  const { loading, data, error } = useQuery<SingleProject>(GET_PROJECT_ON_SLUG, {
    variables: {
      slug,
    },
  })

  if (loading) return <div>Loading...</div>
  if (error) return <div>404 not found</div>

  return (
    <div>
      <div>{data?.project.name}</div>
      <div>{data?.project.desc}</div>
      <div>{data?.project.price}</div>
      <div>{data?.project.deadline}</div>
      <div>{data?.project.created_at}</div>
      <div>{data?.project.updated_at}</div>
      <div>{data?.project.creator.username}</div>
      <button>Bid</button>
    </div>
  )
}

export default ProjectDetails
