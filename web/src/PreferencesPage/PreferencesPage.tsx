import {Button, Container, Form} from "react-bootstrap";
import React, {useState} from "react";

export function PreferencesPage() {
    const options = ["hello", "world"]

    return <Container>
        <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label>Список групп</Form.Label>
                <Form.Control as="select" multiple>
                </Form.Control>
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicPassword">
                <Form.Label>Password</Form.Label>
                <Form.Control type="password" placeholder="Password" />
            </Form.Group>
            <Form.Group className="mb-3" controlId="formBasicCheckbox">
                <Form.Check type="checkbox" label="Check me out" />
            </Form.Group>
            <Button variant="primary">
                Submit
            </Button>
        </Form>
    </Container>
}