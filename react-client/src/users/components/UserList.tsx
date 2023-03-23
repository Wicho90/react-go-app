import { useEffect, useState } from 'react';

interface user {
  id: string;
  name: string;
}
const initialState: user[] =[{
  id: '1',
  name: 'Luis'
}]

export const UserList = () => {

  const [users, setUsers] = useState<user[]>(initialState);

  
  useEffect(() => {
    async function loadUsers() {
      const response: Response = await fetch('http://localhost:8080/users');
      const data = await response.json();
      setUsers(data.data);
      users.forEach(user => console.log(user));
    }
    loadUsers()
  }, [])


  return (
    <ul>
      {
        users.map(user => <li key={user.id}>{user.name}</li>)
      }
    </ul>
  )
}
