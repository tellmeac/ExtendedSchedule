import {Button, Container, Form} from "react-bootstrap";
import React, {useEffect, useState} from "react";
import {FacultyInfo, GroupInfo} from "../Shared/Models";
import {UserConfig} from "./Models";
import {getUserConfig} from "./Api";
import log from "loglevel";

export function SettingsPage() {
    const [userConfig, setUserConfig] = useState<UserConfig>()
    const [isChanged, setIsChanged] = useState<boolean>(false)

    useEffect(()=>{
        getUserConfig().then(config => {
            setUserConfig(config)
        }).catch(err => {
            log.error(err)
        })
    }, [])

    /**
     * План для страницы с настройками:
     * 1. base schedule (open modal to select group)
     * 2. extended group list (with button to open modal to select group to add to list)
     */

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