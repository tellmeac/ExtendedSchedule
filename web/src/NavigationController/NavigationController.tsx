import React from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationController.css"
import GoogleLogin, {GoogleLoginResponse, GoogleLoginResponseOffline, GoogleLogout} from "react-google-login";
import {authProps} from "../auth";
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {resetUserData, selectLoginResponse, updateUserData} from "../Shared/Store/UserSlice";
import {NavUserInfo} from "./NavUserInfo";
import {getUserAuthContentFromResponse} from "../Shared/Models/Auth";

const Title = "Расписание"

export function NavigationController() {
    const userData = useAppSelector(selectLoginResponse)
    const dispatch = useAppDispatch()

    const responseGoogle = (response: (GoogleLoginResponse | GoogleLoginResponseOffline)) => {
        const r = response as GoogleLoginResponse;
        dispatch(updateUserData(getUserAuthContentFromResponse(r)))
    }

    const logout = () => {
        dispatch(resetUserData())
    }

    return <Navbar bg="light" expand="lg">
        <Navbar.Brand className="title" href="/">{Title}</Navbar.Brand>
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
                            onSuccess={responseGoogle}
                            onFailure={err => console.log('fail', err)}
                            isSignedIn={true}
                            cookiePolicy={'single_host_origin'}
                        >
                            Вход
                        </GoogleLogin>
                    </Nav.Item>
                }
                {userData &&
                    <NavUserInfo data={userData} renderLogoutButton={
                        () => <GoogleLogout onLogoutSuccess={logout} clientId={authProps.clientId}/>
                    }/>
                }
            </Nav>
        </Navbar.Collapse>
    </Navbar>
}