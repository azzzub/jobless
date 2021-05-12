import { gql, useMutation } from '@apollo/client'
import { FormEvent, useEffect, useState } from 'react'

const BID_MUTATION = gql`
  mutation Bid($project_id: Int!, $price: Int!, $comment: String!) {
    createBid(input: { project_id: $project_id, price: $price, comment: $comment }) {
      ID
      bidder_id
      project_id
      price
      comment
      created_at
    }
  }
`

interface BiddingProps {
  project_id: number
}

const Bidding: React.FC<BiddingProps> = ({ project_id }) => {
  const [price, setPrice] = useState(0)
  const [comment, setComment] = useState('')
  const [response, setResponse] = useState<Bid>()

  const [bid, { loading, error }] = useMutation<SingleBid>(BID_MUTATION)
  const bidHandler = async (e: FormEvent): Promise<void> => {
    e.preventDefault()
    try {
      const { data } = await bid({
        variables: {
          project_id,
          price,
          comment,
        },
        context: {
          headers: {
            authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        },
      })
      setResponse(data?.bid)
    } catch (error) {
      // TODO
    }
  }

  useEffect(() => {
    // TODO
  }, [response])

  return (
    <div>
      <form onSubmit={bidHandler}>
        <input
          type="number"
          placeholder="Tawaran"
          onChange={(e) => setPrice(parseInt(e.target.value))}
        />
        <input
          type="paragraph"
          placeholder="Keterangan"
          onChange={(e) => setComment(e.target.value)}
        />
        <button type="submit" disabled={loading}>
          Masukan Penawaran
        </button>
      </form>
      {error ? <div>Error</div> : null}
    </div>
  )
}

export default Bidding
