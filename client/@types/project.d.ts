interface Project {
  ID: number
  slug: string
  creator_id: number
  creator: User
  bids: Bid[]
  name: string
  desc: string
  price: number
  deadline: string
  created_at: string
  updated_at: string
  deleted_at: string
}

interface Projects {
  projects: Project[]
}

interface SingleProject {
  project: Project
}
