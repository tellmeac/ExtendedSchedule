import React from "react";
import {format} from "date-fns";
import {useAppDispatch, useAppSelector} from "../../Shared/Hooks";
import {selectSelectedWeekEnd, selectSelectedWeekStart, setPreviousWeek, setNextWeek} from "../../Shared/Store";
import {Button, Container} from "react-bootstrap";
import "./ScheduleController.css"


export const ScheduleController: React.FC = () => {
    const start = useAppSelector(selectSelectedWeekStart)
    const end = useAppSelector(selectSelectedWeekEnd)

    const dispatch = useAppDispatch()

    const setDownWeek = () => {
        dispatch(setPreviousWeek())
    }

    const setUpWeek = () => {
        dispatch(setNextWeek())
    }

    const startFormatted = format(start, "d MMMM u")
    const endFormatted = format(end, "d MMMM u")

    return <Container>
        <Container className={"week-info"}>
            <span>Неделя: </span>
            <span>{startFormatted} - {endFormatted} </span>
        </Container>

        <Container className={"buttons-container"}>
            <Button variant="outline-primary" onClick={setDownWeek}>Предыдущая неделя</Button>
            <Button variant="outline-primary" onClick={setUpWeek}>Следующая неделя</Button>
        </Container>
    </Container>
}