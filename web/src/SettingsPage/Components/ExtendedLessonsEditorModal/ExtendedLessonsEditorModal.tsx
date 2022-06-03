import React, {useEffect, useState} from "react";
import {Button, Form, Modal} from "react-bootstrap";
import {ExtendedLessons, LessonInfo} from "../../Models";
import {LessonItem} from "../LessonInfo";
import "./ExtendedLessonsEditorModal.css"
import {getLessonsInfo} from "../../Api";
import log from "loglevel";

interface Props {
    isOpen: boolean,
    extendedLessons: ExtendedLessons
    selectExtendedLessonsCallback: (lessonsIds: string[]) => void
}

export const ExtendedLessonsEditorModal: React.FC<Props> = ({isOpen, extendedLessons, selectExtendedLessonsCallback}) => {
    const [lessons, setLessons] = useState<LessonInfo[]>([])
    const [selectedLessons, setSelectedLessons] = useState<string[]>(extendedLessons.lessonIds)

    useEffect(()=>{
        if (!isOpen) {
            return
        }

        getLessonsInfo(extendedLessons.group.id).then(info => {
            setLessons(info)
        }).catch(err => {
            log.error(err)
        })
    }, [isOpen])

    const handleClose = () => {
        selectExtendedLessonsCallback(selectedLessons)
    }

    return <Modal show={isOpen} onHide={handleClose}>
        <Modal.Header closeButton>
            <Modal.Title>Выбор дополнительных предметов</Modal.Title>
        </Modal.Header>
        <Modal.Body>
            <Form className="lesson-form">
                {
                    lessons.map((lesson)=>{
                        return <Form.Check key={lesson.id} type="switch"
                                           defaultChecked={selectedLessons.includes(lesson.id)}
                                           label={
                            <div className={"lesson-label"}>
                                <LessonItem lesson={lesson}/>
                            </div>
                        }/>
                    })
                }
            </Form>
        </Modal.Body>
        <Modal.Footer>
            <Button variant="primary" onClick={handleClose}>
                Выбрать
            </Button>
        </Modal.Footer>
    </Modal>
}