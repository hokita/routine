import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

function App() {
  const [tasks, setTasks] = useState([])
  const [task, setTask] = useState('')

  useEffect(() => {
    const fetchData = async ()  => {
      const result = await axios.get(`http://localhost:8082/routines/today/`)
      setTasks(result.data.Tasks)
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

  const handleChangeCheck = (event) => {
    const newTasks = tasks.map((task) => {
      const {done, ...rest} = task
      const targetId = Number(event.target.id)

      if(task.id === targetId) {
        const params = JSON.stringify({ done: event.target.checked });
        axios.put(`http://localhost:8082/tasks/${task.id}/`, params)
          .then(response => {})

        return ({
          done: event.target.checked,
          ...rest
        })
      }

      return ({
        done: task.done,
        ...rest
      })
    })

    setTasks(newTasks)
  }

  const handleSubmit = (event) => {
    event.preventDefault();

    const params = JSON.stringify({ name: task });
    axios.post(`http://localhost:8082/routines/today/`, params)
      .then(response => {
        setTasks(response.data.Tasks)
      })
  }

  const handleDelete = (id, index) => {
    axios.delete(`http://localhost:8082/tasks/${id}/`)
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

  const checkBox = (task) => (
    <input
      id={task.id}
      type="checkbox"
      name="taskCheckbox"
      value={task.name}
      checked={task.done}
      onChange={handleChangeCheck}
    />
  )

  const list = (
    <div>
      <ul>
        { tasks.map((task, index) => {
          return (
              <li key={task.id}>
                { checkBox(task) }
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
