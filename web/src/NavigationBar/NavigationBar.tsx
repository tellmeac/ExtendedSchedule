import React, {useState} from "react";
import {Container, Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationBar.css"
import {Link} from "react-router-dom";
import {GoogleLogin} from "@react-oauth/google";
import jwtDecode from "jwt-decode";
import {storeUserJwtToken} from "../Shared/Api/Token";
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {selectUserInfo, selectSignedIn, setCredentials} from "../Shared/Store";

/**
 * Main navigation bar. Contains user menu and navigation links
 * @constructor
 */
export function NavigationBar() {
    const dispatch = useAppDispatch()
    const isAuthorized = useAppSelector(selectSignedIn)
    const userInfo = useAppSelector(selectUserInfo)

    // @ts-ignore (not exported response type)
    const onSuccessLogin = (credentialResponse) => {
        // extract username from jwt token
        const claims = jwtDecode<{name: string, picture: string}>(credentialResponse.credential)
        dispatch(setCredentials({
            token: credentialResponse.credential,
            username: claims.name,
            avatarUrl: claims.picture,
        }))
        storeUserJwtToken(credentialResponse.credential)
    }

    return <Navbar bg="light">
        <Container fluid>
            <Navbar.Brand className="title" href="/">TSU Schedule</Navbar.Brand>
            <Navbar.Toggle aria-controls="navbarScroll" />
            <Navbar.Collapse id="navbarScroll">
                <Nav className="me-auto my-2 my-lg-0" navbarScroll>
                    <Nav.Item><Link className={"nav-link"} to="/schedule"><i className="bi bi-calendar-heart"/> Расписание</Link></Nav.Item>
                    <Nav.Item><Link className={"nav-link"} to="/settings"><i className="bi bi-gear"/> Параметры</Link></Nav.Item>
                </Nav>
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
                    <Nav.Item className={"user-context"}>Вы вошли как {userInfo.username}</Nav.Item>
                }
            </Navbar.Collapse>
        </Container>
    </Navbar>
}