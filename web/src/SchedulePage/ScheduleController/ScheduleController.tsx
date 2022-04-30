import React from "react";
import {format} from "date-fns";
import {useAppDispatch, useAppSelector} from "../../Shared/Hooks";
import {selectSelectedWeekEnd, selectSelectedWeekStart, setPreviousWeek} from "../../Shared/Store";
import {Button, Container} from "react-bootstrap";
import "./ScheduleController.css"


export const ScheduleController: React.FC = () => {
    const start = useAppSelector(selectSelectedWeekStart)
    const end = useAppSelector(selectSelectedWeekEnd)

    const dispatch = useAppDispatch()

    const setPrevWeek = () => {
        dispatch(setPreviousWeek())
    }

    const setNextWeek = () => {
        dispatch(setPreviousWeek())
    }

    const startFormatted = format(start, "d MMMM u")
    const endFormatted = format(end, "d MMMM u")

    return <Container>
        <Container className={"week-info"}>
            <span>Неделя: </span>
            <span>{startFormatted} - {endFormatted} </span>
        </Container>

        <Container className={"buttons-container"}>
            <Button variant="outline-primary" className={"prev-button"} onClick={setPrevWeek}>Предыдущаяя неделя</Button>
            <Button variant="outline-primary" className={"next-button"} onClick={setNextWeek}>Следующая неделя</Button>
        </Container>
    </Container>
}