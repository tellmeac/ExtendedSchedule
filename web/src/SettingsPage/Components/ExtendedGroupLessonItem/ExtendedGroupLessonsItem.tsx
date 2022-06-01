import React from "react";
import {ExtendedLessons} from "../../Models";
import "./ExtendedGroupLessonItem.css"
import {Badge, Button, ButtonGroup} from "react-bootstrap";

interface Props {
    isNew: boolean
    data: ExtendedLessons
    editCallback: () => void
    removeCallback: () => void
}

export const ExtendedGroupLessonItem: React.FC<Props> = ({isNew, data, editCallback, removeCallback}) => {
    return <div className="container">
        <div className="item-content">
            <span><b>Группа:</b> {data.group.name} <Badge bg="warning">Новый</Badge></span>
            <div>
                <b>Предметов:</b> {data.lessonIds.length}
            </div>
        </div>
        <ButtonGroup className="controller">
            <Button className="edit-btn" onClick={editCallback}>
                <i className="bi bi-gear"/>
            </Button>
            <Button className="remove-btn" variant={"danger"} onClick={removeCallback}>
                <i className="bi bi-trash3"/>
            </Button>
        </ButtonGroup>
    </div>
}