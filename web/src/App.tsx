import React from 'react';
import {Navigate, Outlet, Route, Routes} from "react-router-dom";
import {SchedulePage} from "./SchedulePage";
import {PreferencesPage} from "./PreferencesPage";
import {NavigationBar} from "./NavigationBar";
import {GoogleOAuthProvider} from "@react-oauth/google";

function App() {

  return <GoogleOAuthProvider clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID || ""}>
    <div className={"app-container"}>
      <NavigationBar/>

      <Routes>
        <Route path="/" element={<Navigate to="schedule"/>}/>
        <Route path="/schedule" element={<SchedulePage/>}/>
        <Route path="/preferences" element={<PreferencesPage/>}/>

        <Route path="*" element={<p>Not found!</p>}/>
      </Routes>

      <Outlet/>
    </div>
  </GoogleOAuthProvider>
}

export default App;
