import React from "react";
import {format} from "date-fns";
import {Button, Container} from "react-bootstrap";
import "./ScheduleControlTab.css"

type Props = {
    startDay: number,
    endDay: number,
    moveWeek: (isForward: boolean) => void
}

/**
 * Schedule week control tab
 * @param startDay
 * @param endDay
 * @param moveWeek
 * @constructor
 */
export const ScheduleControlTab: React.FC<Props> = ({startDay, endDay, moveWeek}) => {
    const startFormatted = format(startDay, "d MMMM u")
    const endFormatted = format(endDay, "d MMMM u")

    return <Container>
        <Container className={"week-info"}>
            <span>Неделя: </span>
            <span>{startFormatted} - {endFormatted} </span>
        </Container>

        <Container className={"buttons-container"}>
            <Button variant="outline-primary" onClick={()=>{moveWeek(false)}}><i className="bi bi-caret-left-square"/> Предыдущая неделя</Button>
            <Button variant="outline-primary" onClick={()=>{moveWeek(true)}}>Следующая неделя <i className="bi bi-caret-right-square"/></Button>
        </Container>
    </Container>
}