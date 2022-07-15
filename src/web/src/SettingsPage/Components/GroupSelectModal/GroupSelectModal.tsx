import {FacultyInfo, GroupInfo} from "../../../Shared/Models";
import React, {ChangeEvent, useEffect, useState} from "react";
import {Button, Form, Modal} from "react-bootstrap";
import {getAllFaculties, getFacultyGroups} from "../../Api";
import log from "loglevel";

interface Props {
    isOpen: boolean
    selectGroupCallback: (group: GroupInfo | undefined) => void
}

export const GroupSelectModal: React.FC<Props> = ({isOpen, selectGroupCallback}) => {
    const [allFaculties, setAllFaculties] = useState<FacultyInfo[]>([])
    const [faculty, setFaculty] = useState<FacultyInfo | undefined>(undefined)
    const [allGroups, setAllGroups] = useState<GroupInfo[]>([])
    const [group, setGroup] = useState<GroupInfo | undefined>(undefined)

    useEffect(()=>{
        if (!isOpen) {
            return
        }

        getAllFaculties().then((faculties)=>{
            setAllFaculties(faculties)
        }).catch(err=>{
            log.error(err)
        })
    }, [isOpen])

    useEffect(()=>{
        if (!isOpen || !faculty) {
            return
        }

        getFacultyGroups(faculty?.id || "").then((groups)=>{
            setAllGroups(groups)
        }).catch(err=>{
            log.error(err)
        })
    }, [isOpen, faculty])

    const handleClose = () => {
        selectGroupCallback(group)
    }

    const handleSelectFaculty = (event: ChangeEvent<HTMLSelectElement>) => {
        setFaculty(allFaculties[parseInt(event.target.value)])
    }

    const handleSelectGroup = (event: ChangeEvent<HTMLSelectElement>) => {
        setGroup(allGroups[parseInt(event.target.value)])
    }

    return <Modal show={isOpen} onHide={handleClose}>
        <Modal.Header closeButton>
            <Modal.Title>Выбор группы</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            <Form>
                <Form.Select aria-label="Факультет" onChange={handleSelectFaculty} className="mb-3">
                    {
                        allFaculties.map((faculty, ind) => {
                            return <option key={ind} value={ind}>{faculty.name}</option>
                        })
                    }
                </Form.Select>
                <Form.Select aria-label="Группа" onChange={handleSelectGroup} className="mb-3">
                    {
                        allGroups.map((group, ind) => {
                            return <option key={ind} value={ind}>{group.name}</option>
                        })
                    }
                </Form.Select>
            </Form>
        </Modal.Body>
        <Modal.Footer>
            <Button variant="primary" onClick={handleClose}>
                Выбрать
            </Button>
        </Modal.Footer>
    </Modal>
}