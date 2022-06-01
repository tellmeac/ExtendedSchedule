import {Button, Container, Form, InputGroup, ListGroup} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";
import {UserConfig} from "./Models";
import {getUserConfig} from "./Api";
import log from "loglevel";
import {ExtendedGroupLessonItem} from "./Components/ExtendedGroupLessonItem";

export function SettingsPage() {
    const [userConfig, setUserConfig] = useState<UserConfig | undefined>()
    const [isChanged, setIsChanged] = useState<boolean>(false)

    useEffect(()=>{
        setUserConfig({
            email: "tellmeac@gmail.com",
            excludedLessons: [],
            extendedGroupLessons: [
                {
                    group: {
                        id: "1",
                        name: "931902"
                    },
                    lessonIds: ["1", "2", "3"]
                },
                {
                    group: {
                        id: "2",
                        name: "931903"
                    },
                    lessonIds: ["2", "4"]
                }
            ],
            id: "1",
            baseGroup: {
                id: "3",
                name: "931901"
            }
        })
        // getUserConfig().then(config => {
        //     setUserConfig(config)
        // }).catch(err => {
        //     log.error(err)
        // })
    }, [])

    /**
     * План для страницы с настройками:
     * 1. base schedule (open modal to select group)
     * 2. extended group list (with button to open modal to select group to add to list)
     */

    return <Container>
        <Form className="w-50 mx-auto">
            <Form.Group className="mb-3">
                <Form.Label>Почта</Form.Label>
                <Form.Control placeholder={userConfig?.email || "undefined"} disabled />
            </Form.Group>
            <Form.Group className="mb-3">
                <Form.Label>Основная группа</Form.Label>
                <InputGroup className="mb-3">
                    <Form.Control placeholder={userConfig?.baseGroup.name || "группа не выбрана"} disabled />
                    <Button><i className="bi bi-gear"/> Сменить</Button>
                </InputGroup>
            </Form.Group>
            <Form.Group className="mb-3">
                <Form.Label>Дополнительные предметы</Form.Label>
                <ListGroup>
                    {userConfig &&
                        userConfig?.extendedGroupLessons.map((extendedLessons)=>{
                            return <ListGroup.Item key={extendedLessons.group.id}>
                                <ExtendedGroupLessonItem isNew={true} data={extendedLessons}
                                                         editCallback={()=>{}}
                                                         removeCallback={()=>{}}
                                />
                            </ListGroup.Item>
                        })
                    }
                </ListGroup>
            </Form.Group>
            <Button variant="success">
                Сохранить настройки
            </Button>
        </Form>
    </Container>
}
