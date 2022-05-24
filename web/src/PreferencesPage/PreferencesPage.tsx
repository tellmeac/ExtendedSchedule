import {Button, Container, Form, ListGroup} from "react-bootstrap";
import React, {ChangeEvent, useEffect, useMemo, useState} from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";
import {useAppSelector} from "../Shared/Hooks";
import {getAllFaculties, getFacultyGroups} from "./Api";

export function PreferencesPage() {

    return <Container>
        <Form>
            {/*<Form.Group className="mb-3" controlId="formBasicEmail">*/}
            {/*    <Form.Label className="mb-3">Добавить группу в список</Form.Label>*/}
            {/*    <Form.Group className="mb-3">*/}
            {/*        <Form.Label>Факультет</Form.Label>*/}
            {/*        <Form.Select aria-label='факультеты' onChange={selectFaculty}>*/}
            {/*            {*/}
            {/*                faculties.map((faculty)=>{*/}
            {/*                    return <option key={faculty.name} value={faculty.id + "__" + faculty.name}>*/}
            {/*                        {faculty.name}*/}
            {/*                    </option>*/}
            {/*                })*/}
            {/*            }*/}
            {/*        </Form.Select>*/}
            {/*        <Form.Label>Группа</Form.Label>*/}
            {/*        <Form.Select aria-label='группы' onChange={selectGroup}>*/}
            {/*            {*/}
            {/*                facultyGroups.map((g)=>{*/}
            {/*                    return <option key={g.name} value={g.id + "__" + g.name}>*/}
            {/*                        {g.name}*/}
            {/*                    </option>*/}
            {/*                })*/}
            {/*            }*/}
            {/*        </Form.Select>*/}
            {/*    </Form.Group>*/}
            {/*    <Button variant="outline-primary" onClick={addGroup} disabled={addGroupActionAllowed}>*/}
            {/*        Добавить*/}
            {/*    </Button>*/}
            {/*</Form.Group>*/}
            {/*<Form.Group className="mb-3">*/}
            {/*    <Form.Label>Список выбранных групп:</Form.Label>*/}
            {/*    <ListGroup>*/}
            {/*        {*/}
            {/*            userGroups.map((g, ind)=>{*/}
            {/*                return <ListGroup.Item key={g.name}>*/}
            {/*                    {ind+1}. {g.name}*/}
            {/*                </ListGroup.Item>*/}
            {/*            })*/}
            {/*        }*/}
            {/*    </ListGroup>*/}
            {/*</Form.Group>*/}
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