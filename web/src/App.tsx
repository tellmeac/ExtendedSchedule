import React from 'react';
import {Navigate, Outlet, Route, Routes} from "react-router-dom";
import {SchedulePage} from "./SchedulePage";
import {PreferencesPage} from "./PreferencesPage";
import {NavigationController} from "./NavigationController";

function App() {

  return <div className={"app-container"}>
    <NavigationController/>

    <Routes>
      <Route path="/" element={<Navigate to="schedule"/>}/>
      <Route path="/schedule" element={<SchedulePage/>}/>
      <Route path="/preferences" element={<PreferencesPage/>}/>

      <Route path="*" element={<p>Not found!</p>}/>
    </Routes>

    <Outlet/>
  </div>
}

export default App;
