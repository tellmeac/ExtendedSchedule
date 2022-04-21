import {Container} from "react-bootstrap";
import React from "react";
import {MockScheduleWeek} from "../Mocks/MockScheduleData";
import {WeekScheduleTable} from "../ScheduleTable";
import {ScheduleController} from "../ScheduleController";

export function ScheduleContainer() {
    return <Container>
        <ScheduleController/>
        <WeekScheduleTable dateStart={new Date()} dateEnd={new Date()} days={MockScheduleWeek}/>
    </Container>
}