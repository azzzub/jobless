import { gql, useMutation } from '@apollo/client'
import { FormEvent, useState } from 'react'

const CREATE_PROJECT_MUTATION = gql`
  mutation CreateProject($name: String!, $price: Int!, $desc: String!, $deadline: String!) {
    createProject(input: { name: $name, price: $price, desc: $desc, deadline: $deadline }) {
      ID
      name
      desc
      price
      deadline
      created_at
    }
  }
`

interface CreateProjectProps {
  callback: () => void
}

const CreateProject: React.FC<CreateProjectProps> = ({ ...props }) => {
  const [name, setName] = useState<string>()
  const [price, setPrice] = useState<number>()
  const [desc, setDesc] = useState<string>()
  const [deadline, setDeadline] = useState<string>()
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  //   const [response, setResponse] = useState<Project>()
  const [createProject, { loading, error }] = useMutation<SingleProject>(CREATE_PROJECT_MUTATION)

  const createHandler = async (e: FormEvent): Promise<void> => {
    e.preventDefault()
    try {
      //   const { data } = await createProject({
      await createProject({
        variables: {
          name,
          price,
          desc,
          deadline,
        },
        context: {
          headers: {
            authorization: `Bearer ${localStorage.getItem('token')}`,
          },
        },
      })
      //   setResponse(data?.project)
      props.callback()
    } catch (error) {
      // TODO
    }
  }

  if (loading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>

  return (
    <div>
      <div>Create Project</div>
      <form onSubmit={createHandler}>
        <input
          type="text"
          placeholder="Project name"
          onChange={(e) => {
            setName(e.target.value)
          }}
        />
        <input
          type="text"
          placeholder="Deskripsi"
          onChange={(e) => {
            setDesc(e.target.value)
          }}
        />
        <input
          type="number"
          placeholder="Harga"
          onChange={(e) => {
            setPrice(parseInt(e.target.value))
          }}
        />
        <input
          type="date"
          placeholder="Deadline"
          onChange={(e) => {
            setDeadline(e.target.value)
          }}
        />
        <button type="submit">Buat proyek</button>
      </form>
    </div>
  )
}

export default CreateProject
