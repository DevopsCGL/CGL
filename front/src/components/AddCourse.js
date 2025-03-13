import React, { useState } from 'react';

function AddCourse({ addCourse }) {
  const [newItem, setNewItem] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (newItem.trim()) {
      addCourse(newItem);
      setNewItem('');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={newItem}
        onChange={(e) => setNewItem(e.target.value)}
        placeholder="Ajouter un article"
      />
      <button type="submit">Ajouter</button>
    </form>
  );
}

export default AddCourse;
