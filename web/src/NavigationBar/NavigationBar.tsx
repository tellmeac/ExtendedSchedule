import React, {useState} from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationBar.css"
import {Link} from "react-router-dom";
import {GoogleLogin} from "@react-oauth/google";
import jwtDecode from "jwt-decode";
import {storeUserJwtToken} from "../Shared/Api/Token";

/**
 * Main navigation bar. Contains user menu and navigation links
 * @constructor
 */
export function NavigationBar() {
    const [isAuthorized, setIsAuthorized] = useState<boolean>(false)
    const [userName, setUserName] = useState<string>("")

    // @ts-ignore (not exported response type)
    const onSuccessLogin = (credentialResponse) => {
        setIsAuthorized(true)

        // extract username from jwt token
        const user = jwtDecode<{name: string}>(credentialResponse.credential)
        setUserName(user.name)
        storeUserJwtToken(credentialResponse.credential)

        console.log(credentialResponse);
    }

    return <Navbar bg="light" expand="lg">
        <Navbar.Brand className="title" href="/">TSU Schedule</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Item><Link className={"nav-link"} to="/schedule"><i className="bi bi-calendar-heart"/> Расписание</Link></Nav.Item>
                <Nav.Item><Link className={"nav-link"} to="/settings"><i className="bi bi-gear"/> Параметры</Link></Nav.Item>
            </Nav>
            <Nav className="mr-auto">
                {
                    !isAuthorized &&
                    <GoogleLogin
                        auto_select
                        useOneTap
                        shape="circle"
                        theme="outline"
                        onSuccess={onSuccessLogin}
                        onError={() => {
                            console.error('Login Failed');
                        }}
                    />
                }
                {
                    // TODO: replace with user info component
                    isAuthorized &&
                    <Nav.Item className={"user-context"}>Вы вошли как {userName}</Nav.Item>
                }
            </Nav>
        </Navbar.Collapse>
    </Navbar>
}