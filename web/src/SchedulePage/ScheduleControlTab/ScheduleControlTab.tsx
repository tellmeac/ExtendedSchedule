import React from "react";
import {format} from "date-fns";
import {useAppDispatch, useAppSelector} from "../../Shared/Hooks";
import {setPreviousWeek, setNextWeek, selectWeekPeriod} from "../../Shared/Store";
import {Button, Container} from "react-bootstrap";
import "./ScheduleControlTab.css"


export const ScheduleControlTab: React.FC = () => {
    const period = useAppSelector(selectWeekPeriod)

    const dispatch = useAppDispatch()

    const setDownWeek = () => {
        dispatch(setPreviousWeek())
    }

    const setUpWeek = () => {
        dispatch(setNextWeek())
    }

    const startFormatted = format(period.weekStart, "d MMMM u")
    const endFormatted = format(period.weekEnd, "d MMMM u")

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