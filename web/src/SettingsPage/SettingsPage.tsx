import {Button, Container, Form} from "react-bootstrap";
import React from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";

export function SettingsPage() {

    return <Container>
        <Form>
            <Button variant="success">
                Сохранить
            </Button>
        </Form>
    </Container>
}

function parseFaculty(value: string): FacultyInfo {
    const [id, name] = value.split("__", 2)
    return {
        id: id,
        name: name
    }
}

function parseGroup(value: string): GroupInfo {
    const [id, name] = value.split("__", 2)
    return {
        id: id,
        name: name
    }
}