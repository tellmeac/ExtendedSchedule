import React from "react";
import {WeekSchedule} from "./WeekSchedule";
import {ScheduleController} from "./ScheduleController/ScheduleController";
import {Container} from "react-bootstrap";
import "./SchedulePage.css"

export function SchedulePage() {
    return <Container className={"schedule-page"}>
        <ScheduleController/>
        <WeekSchedule/>
    </Container>
}