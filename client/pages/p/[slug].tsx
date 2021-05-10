import { useRouter } from 'next/router'

const ProjectDetails: React.FC = () => {
  const router = useRouter()
  const { slug } = router.query

  return <div>{<div>{slug}</div>}</div>
}

export default ProjectDetails
