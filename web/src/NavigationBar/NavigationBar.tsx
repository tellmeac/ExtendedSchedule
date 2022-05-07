import React from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationBar.css"
import GoogleLogin, {GoogleLoginResponse, GoogleLoginResponseOffline, GoogleLogout} from "react-google-login";
import {authProps} from "../auth";
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {selectLoginResponse, updateUserData} from "../Shared/Store";
import {getUserAuthContentFromResponse} from "../Shared/Models/Auth";
import {UserMenu} from "./UserMenu";
import {useNavigate} from "react-router-dom";

export function NavigationBar() {
    const navigate = useNavigate()

    const dispatch = useAppDispatch()
    const userData = useAppSelector(selectLoginResponse)

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
                <Nav.Link href="/schedule">Расписание</Nav.Link>
                <Nav.Link href="/preferences">Параметры</Nav.Link>
            </Nav>
            <Nav className="mr-auto">
                {!userData &&
                    <Nav.Item>
                        <GoogleLogin
                            clientId={authProps.clientId}
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
                    <UserMenu data={userData} renderLogoutButton={
                        () => <GoogleLogout onLogoutSuccess={logoutSuccess} clientId={authProps.clientId}/>
                    }/>
                }
            </Nav>
        </Navbar.Collapse>
    </Navbar>
}