import React from "react";
import {WeekSchedule} from "./WeekSchedule";
import {ScheduleControlTab} from "./ScheduleControlTab/ScheduleControlTab";
import {Container} from "react-bootstrap";
import "./SchedulePage.css"

export function SchedulePage() {
    return <Container className={"schedule-page"}>
        <ScheduleControlTab/>
        <WeekSchedule/>
    </Container>
}