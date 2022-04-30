import React from "react";
import {WeekSchedule} from "./WeekSchedule";
import {ScheduleController} from "./ScheduleController/ScheduleController";
import {Container} from "react-bootstrap";

export function SchedulePage() {
    return <Container>
        <ScheduleController/>
        <WeekSchedule/>
    </Container>
}