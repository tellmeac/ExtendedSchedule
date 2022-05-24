import {Button, Container, Form, ListGroup} from "react-bootstrap";
import React, {ChangeEvent, useEffect, useMemo, useState} from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";
import {useAppSelector} from "../Shared/Hooks";
import {selectUserData} from "../Shared/Store";
import {getAllFaculties, getFacultyGroups} from "./Api";


export function PreferencesPage() {
    const user = useAppSelector(selectUserData)
    const [userGroups, setUserGroups] = useState<GroupInfo[]>([])

    // for new group selection
    const [faculties, setFaculties] = useState<FacultyInfo[]>([])
    const [facultyGroups, setFacultyGroups] = useState<GroupInfo[]>([])
    const [selectedFaculty, setSelectedFaculty] = useState<FacultyInfo | undefined>(undefined)
    const [selectedGroup, setSelectedGroup] = useState<GroupInfo | undefined>(undefined)

    const [addGroupActionAllowed, setAddGroupActionAllowed] = useState<boolean>(false)

    useEffect(()=>{
        // TODO:
        // get user's joined groups

        setUserGroups([])
    }, [user])

    useEffect(()=>{
        if (user === undefined) {
            return
        }

        getAllFaculties(user?.tokenId || "").then(r => {
            setFaculties(r)
            setSelectedFaculty(r[0])
        }).catch(err => {
            console.error(err)
        })
    }, [])

    useEffect(()=>{
        if (user === undefined) {
            return
        }

        if (selectedFaculty === undefined) {
            return
        }

        getFacultyGroups(user?.tokenId || "", selectedFaculty.id || "").then(r => {
            setFacultyGroups(r)
        }).catch(err => {
            console.error(err)
        })
    }, [user, selectedFaculty])

    const selectFaculty = (event: ChangeEvent<HTMLSelectElement>) => {
        setSelectedFaculty(parseFaculty(event.target.value))
    }

    const selectGroup = (event: ChangeEvent<HTMLSelectElement>) => {
        setSelectedGroup(parseGroup(event.target.value))
        setAddGroupActionAllowed(!isAddGroupActionDisabled())
    }

    const addGroup = () => {
        if (!selectedGroup) {
            return
        }
        setUserGroups([...userGroups, selectedGroup])
        setAddGroupActionAllowed(false)
    }

    const isAddGroupActionDisabled = () => {
        return selectedGroup === undefined || userGroups.find((g)=>{return g === selectedGroup}) !== undefined
    }

    return <Container>
        <Form>
            <Form.Group className="mb-3" controlId="formBasicEmail">
                <Form.Label className="mb-3">Добавить группу в список</Form.Label>
                <Form.Group className="mb-3">
                    <Form.Label>Факультет</Form.Label>
                    <Form.Select aria-label='факультеты' onChange={selectFaculty}>
                        {
                            faculties.map((faculty)=>{
                                return <option key={faculty.name} value={faculty.id + "__" + faculty.name}>
                                    {faculty.name}
                                </option>
                            })
                        }
                    </Form.Select>
                    <Form.Label>Группа</Form.Label>
                    <Form.Select aria-label='группы' onChange={selectGroup}>
                        {
                            facultyGroups.map((g)=>{
                                return <option key={g.name} value={g.id + "__" + g.name}>
                                    {g.name}
                                </option>
                            })
                        }
                    </Form.Select>
                </Form.Group>
                <Button variant="outline-primary" onClick={addGroup} disabled={addGroupActionAllowed}>
                    Добавить
                </Button>
            </Form.Group>
            <Form.Group className="mb-3">
                <Form.Label>Список выбранных групп:</Form.Label>
                <ListGroup>
                    {
                        userGroups.map((g, ind)=>{
                            return <ListGroup.Item key={g.name}>
                                {ind+1}. {g.name}
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