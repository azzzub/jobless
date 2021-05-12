interface Bid {
  ID: number
  bidder_id: number
  bidder: User
  project_id: number
  project: Project
  price: number
  comment: string
  created_at: string
  updated_at: string
  deleted_at: string
}

interface Bids {
  bids: Bid[]
}

interface SingleBid {
  bid: Bid
}
