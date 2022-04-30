import React from "react";
import {ScheduleController} from "./ScheduleController";
import {WeekScheduleTable} from "./ScheduleTable";

export function SchedulePage() {
    return <>
        <ScheduleController/>
        <WeekScheduleTable dateStart={new Date()} dateEnd={new Date()}/>
    </>
}