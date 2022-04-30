import React from "react";
import {Nav, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./NavigationController.css"

const Title = "Расписание"

export function NavigationController() {
    return <Navbar bg="light" expand="lg">
        <Navbar.Brand className="title" href="/">{Title}</Navbar.Brand>
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
                <Nav.Link href="/">Обновить</Nav.Link>
                <Nav.Link href="/preferences">Параметры</Nav.Link>
            </Nav>
        </Navbar.Collapse>
    </Navbar>
}