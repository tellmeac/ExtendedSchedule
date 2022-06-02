import {Button, Container, Form, InputGroup, ListGroup} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import {ExtendedLessons, UserConfig} from "./Models";
import {ExtendedGroupLessonItem} from "./Components/ExtendedGroupLessonItem";
import {GroupSelectModal} from "./Components/GroupSelectModal/GroupSelectModal";

export function SettingsPage() {
    const [userConfig, setUserConfig] = useState<UserConfig | undefined>()
    const [newExtendedGroupIds, setNewExtendedGroupIds] = useState<string[]>([])
    const [isChanged, setIsChanged] = useState<boolean>(false)

    const [selectGroupModal, setSelectGroupModal] = useState<boolean>(false)

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

    const isNewExtendedGroup = (extended: ExtendedLessons) => {
        return newExtendedGroupIds.includes(extended.group.id)
    }

    const addExtendedGroups = () => {
        setSelectGroupModal(true)
    }

    return <>
        <Container>
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
                    <ListGroup className="mb-3">
                        {userConfig &&
                        userConfig?.extendedGroupLessons.map((extendedLessons)=>{
                            return <ListGroup.Item key={extendedLessons.group.id}>
                                <ExtendedGroupLessonItem isNew={isNewExtendedGroup(extendedLessons)}
                                                         data={extendedLessons}
                                                         editCallback={()=>{}}
                                                         removeCallback={()=>{}}
                                />
                            </ListGroup.Item>
                        })
                        }
                    </ListGroup>
                    <Button variant="outline-success" onClick={addExtendedGroups}>Добавить</Button>
                </Form.Group>
                <Button variant="success">
                    Сохранить настройки
                </Button>
            </Form>
        </Container>

        <GroupSelectModal isOpen={selectGroupModal} selectGroupCallback={(x)=>{setSelectGroupModal(false)}}/>
    </>
}
