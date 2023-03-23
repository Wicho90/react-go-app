import { useState, FormEvent, ChangeEvent } from 'react';

export const Form = () => {
  const [name, setName] = useState("");
  const uri = "http://localhost:8080/users"


  const handleSumbmit = async(e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const response = await fetch(uri, {
      method: 'POST',
      body: JSON.stringify({ name }),
      headers: {
        "Content-Type": "application/json"
      }
    })

    const data = await response.json()
    console.log(data);
    
  }

  const onInputChange = ({ target }: ChangeEvent<HTMLInputElement>) => {
    setName(target.value)
  }


  return (
    <form onSubmit={handleSumbmit} autoComplete="off">
        <input 
          type="text"
          placeholder="Coloca tu nombre"
          value={name}
          onChange={onInputChange}  
        />
        <button type="submit">Guardar</button>
      </form>
  )
}
