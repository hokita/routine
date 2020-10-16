import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([])
  const [task, setTask] = useState('')

  useEffect(() => {
    const fetchData = async ()  => {
      const result = await axios.get(`http://localhost:8080/tasks/`)
      setTasks(result.data)
    }

    fetchData()
  }, [])

  const handleChange = (event) => {
    switch(event.target.name) {
      case 'task':
        setTask(event.target.value)
        break
      default:
    }
  }

  const handleSubmit = (event) => {
    event.preventDefault();

    const params = JSON.stringify({ name: task });
    axios.post(`http://localhost:8080/tasks/`, params)
      .then(response => {
        const newTasks = [...tasks]
        newTasks.push(response.data)
        console.log(response.data)
        setTasks(newTasks)
      })
  }

  const handleDelete = (id, index) => {
    axios.delete(`http://localhost:8080/tasks/${id}`)
      .then(response => {
        const newTasks = [...tasks]
        newTasks.splice(index, 1)
        setTasks(newTasks)
      })
  }

  const inputForm = (
    <form onSubmit={handleSubmit}>
      <label>
        new task:
        <input type="text" name="task" value={task} onChange={handleChange} />
      </label>
      <input type="submit" value="Add" />
    </form>
  )

  const list = (
    <div>
      <ul>
        { tasks.map((task, index) => {
          return (
              <li key={task.id}>
                {task.name} <button onClick={() => handleDelete(task.id, index)}>削除</button>
              </li>
          )
        })}
      </ul>
    </div>
  )

  return (
    <div className="App">
      <header className="App-header">
        { inputForm }
        { list }
      </header>
    </div>
  );
}

export default App;
