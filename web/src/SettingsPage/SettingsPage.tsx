import {Button, Container, Form, InputGroup, ListGroup} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import {ExtendedLessons, UserConfig} from "./Models";
import {ExtendedGroupLessonItem} from "./Components/ExtendedGroupLessonItem";
import {GroupSelectModal} from "./Components/GroupSelectModal/GroupSelectModal";
import {GroupInfo} from "../Shared/Models";
import log from "loglevel";
import "./SettingsPage.css"
import {useAppSelector} from "../Shared/Hooks";
import {selectSignedIn} from "../Shared/Store";
import {getUserConfig} from "./Api";
import {ExtendedLessonsEditorModal} from "./Components/ExtendedLessonsEditorModal";

export function SettingsPage() {
    const isAuthorized = useAppSelector(selectSignedIn)

    const [userConfig, setUserConfig] = useState<UserConfig>({
        baseGroup: undefined,
        email: "",
        excludedLessons: [],
        extendedGroupLessons: [],
        id: ""
    })
    const [configExtendedRender, setConfigExtendedRender] = useState<ExtendedLessons[]>([])
    const [configChanged, setConfigChanged] = useState<boolean>(false)

    const [isOpenGroupModal, setOpenGroupModal] = useState<boolean>(false)

    const [isOpenExtendedLessonsEditor, setOpenExtendedLessonsEditor] = useState<boolean>(false)
    const [selectedExtendedGroupToEdit, setSelectedExtendedGroupToEdit] = useState<ExtendedLessons>({
        group: {
            id: "",
            name: "undefined"
        },
        lessonIds: []
    })

    useEffect(()=>{
        if (!isAuthorized) {
            return
        }
        getUserConfig().then(config => {
            setUserConfig(config)
            setConfigExtendedRender(config.extendedGroupLessons)
        }).catch(err => {
            log.error(err)
        })
    }, [isAuthorized])

    const editExtendedGroupCallback = (extended: ExtendedLessons) => {
        setSelectedExtendedGroupToEdit(extended)
        setOpenExtendedLessonsEditor(true)
    }

    const removeExtendedGroupCallback = (extended: ExtendedLessons) => {
        const updated = userConfig
        updated.extendedGroupLessons = userConfig.extendedGroupLessons.filter((el)=>{
            return el.group.id !== extended.group.id
        })
        setUserConfig(updated)
        setConfigExtendedRender(updated.extendedGroupLessons)

        console.log(userConfig)
    }

    const addExtendedGroups = () => {
        setOpenGroupModal(true)
    }

    return <>
        <Container className="settings-container">
            <Form className="mx-auto">
                <Form.Group className="mb-3">
                    <Form.Label>Почта</Form.Label>
                    <Form.Control placeholder={userConfig.email || "undefined"} disabled />
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Основная группа</Form.Label>
                    <InputGroup className="mb-3">
                        <Form.Control placeholder={userConfig.baseGroup?.name || "группа отсутствует"} disabled />
                        <Button><i className="bi bi-gear"/> Сменить</Button>
                    </InputGroup>
                </Form.Group>
                <Form.Group className="mb-3">
                    <Form.Label>Дополнительные предметы</Form.Label>
                    <ListGroup className="mb-3">
                        {
                            configExtendedRender.map((extendedLessons)=>{
                                return <ListGroup.Item key={extendedLessons.group.id}>
                                    <ExtendedGroupLessonItem isNew={false}
                                                             data={extendedLessons}
                                                             editCallback={()=>{
                                                                 editExtendedGroupCallback(extendedLessons)
                                                             }}
                                                             removeCallback={()=>{
                                                                 removeExtendedGroupCallback(extendedLessons)
                                                             }}
                                    />
                                </ListGroup.Item>
                            })
                        }
                    </ListGroup>
                    <Button variant="outline-success" onClick={addExtendedGroups}>Добавить</Button>
                </Form.Group>
                <Button variant="success" disabled={!configChanged}>
                    Сохранить настройки
                </Button>
            </Form>
        </Container>

        {/* TODO: use callbacks to save data to new user config */}
        <GroupSelectModal isOpen={isOpenGroupModal} selectGroupCallback={(x: GroupInfo | undefined)=>{setOpenGroupModal(false)}}/>
        <ExtendedLessonsEditorModal isOpen={isOpenExtendedLessonsEditor} extendedLessons={selectedExtendedGroupToEdit} selectExtendedLessonsCallback={(y: string[]) => {setOpenExtendedLessonsEditor(false)}}/>
    </>
}
