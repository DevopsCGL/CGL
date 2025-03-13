import React from 'react';

function CourseList({ courses }) {
  return (
    <div>
      <ul>
        {courses.map((course) => (
          <li key={course.id}>{course.item}</li>
        ))}
      </ul>
    </div>
  );
}

export default CourseList;
