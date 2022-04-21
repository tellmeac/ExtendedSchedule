import React from "react";
import {Button, Form, Navbar} from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import "./ScheduleController.css"

const Title = "Расписание"

export function ScheduleController() {
    return <Navbar bg="light" expand="lg">
        <Navbar.Brand className="title" href="/">{Title}</Navbar.Brand>
        <Navbar.Toggle aria-controls="navbarScroll" />
        <Navbar.Collapse id="navbarScroll">

        <Form className="me-auto my-2 my-lg-0">
            <Button variant="outline-primary" className="btn-refresh me-2"
            >
                Обновить
            </Button>

            <Button variant="outline-success"
            >
                Параметры
            </Button>
        </Form>

        </Navbar.Collapse>

    </Navbar>
}