import React, { useState, useEffect } from 'react';
import axios from 'axios';
import logo from './logo.svg';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([])

  useEffect(() => {
    const fetchData = async ()  => {
      const result = await axios.get(`http://localhost:8080/tasks/`)
      setTasks(result.data)
    }

    fetchData()
  }, [])

  const handleDelete = (id, index) => {
    axios.delete(`http://localhost:8080/tasks/${id}`)
      .then(response => {
        const newTasks = [...tasks]
        newTasks.splice(index, 1)
        setTasks(newTasks)
      })
  }

  const list = (
    <ul>
      { tasks.map((task, index) => {
        return (
          <li key={task.id}>
            taskname: {task.name} <button onClick={() => handleDelete(task.id, index)}>削除</button>
          </li>
        )
      })}
    </ul>
  )

  return (
    <div className="App">
      <header className="App-header">
        { list }
      </header>
    </div>
  );
}

export default App;
