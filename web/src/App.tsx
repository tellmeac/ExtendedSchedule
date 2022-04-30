import React from 'react';
import './App.css';
import {Navigate, Outlet, Route, Routes} from "react-router-dom";
import {SchedulePage} from "./SchedulePage";
import {PreferencesPage} from "./PreferencesPage";
import {Container} from "react-bootstrap";
import {NavigationController} from "./NavigationController";

function App() {
  return <Container>
    <NavigationController/>

    <Routes>
      <Route path="/" element={<Navigate to="schedule"/>}/>
      <Route path="/schedule" element={<SchedulePage/>}/>
      <Route path="/preferences" element={<PreferencesPage/>}/>

      <Route path="*" element={<p>Not found!</p>}/>
    </Routes>

    <Outlet/>
  </Container>
}

export default App;
