import {Container} from "react-bootstrap";
import React from "react";
import {MockScheduleWeek} from "../Mocks/MockScheduleData";
import {WeekScheduleTable} from "../ScheduleTable";

export function ScheduleContainer() {
    return <Container>
        <WeekScheduleTable dateStart={new Date()} dateEnd={new Date()} days={MockScheduleWeek}/>
    </Container>
}