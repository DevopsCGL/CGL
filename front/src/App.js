import React, { useState, useEffect } from 'react';
import './App.css';
import CourseList from './components/CourseList';
import AddCourse from './components/AddCourse';

function App() {
  const [courses, setCourses] = useState([]);

  // Récupérer la liste des courses depuis l'API au chargement de la page
  useEffect(() => {
    fetch('http://localhost:5000/courses')
      .then((response) => response.json())
      .then((data) => setCourses(data));
  }, []);

  // Ajouter un nouvel élément à la liste
  const addCourse = (newItem) => {
    fetch('http://localhost:5000/courses/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ item: newItem }),
    })
      .then((response) => response.json())
      .then((data) => {
        setCourses([...courses, data]);
      });
  };

  return (
    <div className="App">
      <h1>Liste de Courses</h1>
      <CourseList courses={courses} />
      <AddCourse addCourse={addCourse} />
    </div>
  );
}

export default App;
