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
        console.info(response.data.message)
        tasks.splice(index, 1)
        setTasks(tasks)
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
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload..
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      <ul>
        { list }
      </ul>
      </header>
    </div>
  );
}

export default App;
