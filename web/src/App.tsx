import React from 'react';
import './App.css';
import {Navigate, Outlet, Route, Routes} from "react-router-dom";
import {SchedulePage} from "./SchedulePage";

function App() {
  return <>
    <Routes>
      <Route path="/" element={<Navigate to="schedule"/>}/>
      <Route path="/schedule" element={<SchedulePage/>}/>

      <Route path="*" element={<p>Not found!</p>}/>
    </Routes>

    <Outlet/>
  </>
}

export default App;
