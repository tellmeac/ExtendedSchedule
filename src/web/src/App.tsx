import React from 'react';
import {Outlet, Route, Routes} from "react-router-dom";
import {SchedulePage} from "./SchedulePage";
import {SettingsPage} from "./SettingsPage";
import {NavigationBar} from "./NavigationBar";
import {GoogleOAuthProvider} from "@react-oauth/google";

function App() {
  return <GoogleOAuthProvider clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID || ""}>
    <div>
      <NavigationBar/>

      <Routes>
        <Route path="/" element={<SchedulePage/>}/>
        <Route path="/schedule" element={<SchedulePage/>}/>
        <Route path="/settings" element={<SettingsPage/>}/>

        <Route path="*" element={<p>404. Page not found!</p>}/>
      </Routes>

      <Outlet/>
    </div>
  </GoogleOAuthProvider>
}

export default App;
