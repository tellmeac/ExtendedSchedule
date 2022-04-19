import React from 'react';
import logo from './logo.svg';
import './App.css';
import {WeekScheduleTable} from "./Schedule/ScheduleTable";
import {MockScheduleWeek} from "./Schedule/Mocks/MockScheduleData";

function App() {
  return (
    <div className="App">

      <WeekScheduleTable dateStart={new Date()} dateEnd={new Date()} days={MockScheduleWeek}/>

    </div>
  );
}

export default App;
