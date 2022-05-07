import React from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationBar.css"
import GoogleLogin, {GoogleLoginResponse, GoogleLoginResponseOffline, GoogleLogout} from "react-google-login";
import {authProps} from "../auth";
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {resetUserData, selectLoginResponse, updateUserData, updateSchedule} from "../Shared/Store";
import {getUserAuthContentFromResponse} from "../Shared/Models/Auth";

export function NavigationBar() {
    const userData = useAppSelector(selectLoginResponse)
    const dispatch = useAppDispatch()

    const loginSuccess = (response: (GoogleLoginResponse | GoogleLoginResponseOffline)) => {
        const r = response as GoogleLoginResponse;
        dispatch(updateUserData(getUserAuthContentFromResponse(r)))
    }

    const logoutSuccess = () => {
        dispatch(resetUserData())
        dispatch(updateSchedule([]))
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
                            onFailure={err => console.log('fail', err)}
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