import React from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationBar.css"
import GoogleLogin, {GoogleLoginResponse, GoogleLoginResponseOffline, GoogleLogout} from "react-google-login";
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {selectUserData, updateUserData} from "../Shared/Store";
import {getUserAuthContentFromResponse} from "../Shared/Models/Auth";
import {UserMenu} from "./UserMenu";
import {Link, useNavigate} from "react-router-dom";

export function NavigationBar() {
    const navigate = useNavigate()

    const dispatch = useAppDispatch()
    const userData = useAppSelector(selectUserData)

    const loginSuccess = (response: (GoogleLoginResponse | GoogleLoginResponseOffline)) => {
        const r = response as GoogleLoginResponse;
        dispatch(updateUserData(getUserAuthContentFromResponse(r)))
    }

    const logoutSuccess = () => {
        navigate(0)
    }

    return <Navbar bg="light" expand="lg">
        <Navbar.Brand className="title" href="/">Расписание</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Item><Link className={"nav-link"} to="/schedule">Расписание</Link></Nav.Item>
                <Nav.Item><Link className={"nav-link"} to="/preferences">Параметры</Link></Nav.Item>
            </Nav>
            <Nav className="mr-auto">
                {!userData &&
                    <Nav.Item>
                        <GoogleLogin
                            clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID || ""}
                            onSuccess={loginSuccess}
                            onFailure={err => console.log('failed to sign in', err)}
                            isSignedIn={true}
                            cookiePolicy={'single_host_origin'}
                        >
                            Вход
                        </GoogleLogin>
                    </Nav.Item>
                }
                {userData &&
                    <Nav.Item>
                        <UserMenu data={userData} renderLogoutButton={
                            () => <GoogleLogout onLogoutSuccess={logoutSuccess} clientId={process.env.GOOGLE_CLIENT_ID || ""}/>
                        }/>
                    </Nav.Item>
                }
            </Nav>
        </Navbar.Collapse>
    </Navbar>
}