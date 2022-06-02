import {FacultyInfo, GroupInfo} from "../../../Shared/Models";
import React, {useEffect, useState} from "react";
import {Button, Modal} from "react-bootstrap";
import {getAllFaculties} from "../../Api";
import log from "loglevel";

interface Props {
    isOpen: boolean
    selectGroupCallback: (group: GroupInfo | undefined) => void
}

export const GroupSelectModal: React.FC<Props> = ({isOpen, selectGroupCallback}) => {
    const [allFaculties, setAllFaculties] = useState<FacultyInfo[]>([])
    const [faculty, setFaculty] = useState<FacultyInfo | undefined>(undefined)
    const [group, setGroup] = useState<GroupInfo | undefined>(undefined)

    useEffect(()=>{
        getAllFaculties().then((faculties)=>{
            setAllFaculties(faculties)
        }).catch(err=>{
            log.error(err)
        })
    })

    const handleClose = () => {
        selectGroupCallback(group)
    }

    return <Modal show={isOpen} onHide={handleClose}>
        <Modal.Header closeButton>
            <Modal.Title>Выбор группы</Modal.Title>
        </Modal.Header>
        <Modal.Body>Выберите группу</Modal.Body>
        <Modal.Footer>
            <Button variant="primary" onClick={handleClose}>
                Выбрать
            </Button>
        </Modal.Footer>
    </Modal>
}