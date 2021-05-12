import { gql, useQuery } from '@apollo/client'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useState } from 'react'
import Bidding from '../../components/bidding'

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
      bids {
        bidder {
          username
        }
        price
        comment
        created_at
      }
      created_at
      updated_at
    }
  }
`

const ProjectDetails: React.FC = () => {
  const router = useRouter()
  const [isBidClicked, setIsBidClicked] = useState(false)
  const { slug } = router.query
  const { loading, data, error, refetch } = useQuery<SingleProject>(GET_PROJECT_ON_SLUG, {
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
      {data?.project.bids.map((bid, i) => {
        return (
          <div key={i}>
            <div>{bid.bidder.username}</div>
            <div>{bid.comment}</div>
            <div>{bid.price}</div>
          </div>
        )
      })}
      <button onClick={() => setIsBidClicked(true)}>Bid</button>
      {isBidClicked ? (
        <Bidding
          project_id={data?.project.ID == undefined ? 0 : data.project.ID}
          callback={() => refetch()}
        />
      ) : null}
      <Link href="/">
        <button>Back</button>
      </Link>
    </div>
  )
}

export default ProjectDetails
