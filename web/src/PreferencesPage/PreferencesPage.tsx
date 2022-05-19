import {Button, Container, Form, ListGroup} from "react-bootstrap";
import React, {useState} from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";


export function PreferencesPage() {
    const mockFaculties: FacultyInfo[] = [
        {
            id: "1",
            name: "Прикладная математика и компьютерные науки"
        },
        {
            id: "2",
            name: "Прикладная математика и компьютерные науки"
        }
    ]

    const mockGroups = [
        {
            "id": "931901",
            "name": "11"
        },
        {
            "id": "931902",
            "name": "12"
        },
        {
            "id": "931903",
            "name": "13"
        }
    ]

    const mockInitialGroups = [
        {
            "id": "931901",
            "name": "11"
        },
        {
            "id": "777",
            "name": "14"
        },
        {
            "id": "666",
            "name": "15"
        }
    ]

    const [userGroups, setUserGroups] = useState<GroupInfo[]>(mockInitialGroups)
    const [selectedGroup, setSelectedGroup] = useState<GroupInfo | undefined>(undefined)

    const addGroup = () => {
        if
        setUserGroups([...userGroups, selectedGroup])
    }

    return <Container>
        <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label className="mb-3">Добавить группу в список</Form.Label>
                <Form.Group className="mb-3">
                    <Form.Label>Факультет</Form.Label>
                    <Form.Select aria-label=''>
                        {
                            mockFaculties.map((faculty)=>{
                                return <option id={faculty.value} value={faculty.value}>
                                    {faculty.labelKey}
                                </option>
                            })
                        }
                    </Form.Select>
                    <Form.Label>Группа</Form.Label>
                    <Form.Select aria-label='' onSelect={selectGroup}>
                        {
                            mockGroups.map((g)=>{
                                return <option id={g.value} value={g.value}>
                                    {g.labelKey}
                                </option>
                            })
                        }
                    </Form.Select>
                </Form.Group>
                <Button variant="outline-primary" onClick={addGroup} disabled={selectedGroup === undefined}>
                    Добавить
                </Button>
            </Form.Group>
            <Form.Group className="mb-3">
                <Form.Label>Список выбранных групп:</Form.Label>
                <ListGroup>
                    {
                        userGroups.map((g)=>{
                            return <ListGroup.Item id={g.id}>
                                {g.name}
                            </ListGroup.Item>
                        })
                    }
                </ListGroup>
            </Form.Group>
            <Button variant="success">
                Сохранить
            </Button>
        </Form>
    </Container>
}