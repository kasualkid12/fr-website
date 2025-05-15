import React from 'react';

import '../styles/Schedule.scss'

function Schedule() {
  return (
    <div className='Schedule'>
      <h2>Schedule</h2>
      <p>All times in Eastern</p>
      <div className='Days'>
        <div className='Day'>
        <h2>Friday</h2>
        <p>To be Announced</p>
        </div>
        <div className='Day'>
        <h2>Saturday</h2>
        <p>To be Announced</p>
        </div>
        <div className='Day'>
        <h2>Sunday</h2>
        <p>To be Announced</p>
        </div>
      </div>
    </div>
  )
}

export default Schedule;