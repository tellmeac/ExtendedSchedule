import React from "react";
import {WeekSchedule} from "./WeekSchedule";

export function SchedulePage() {
    return <WeekSchedule dateStart={new Date()} dateEnd={new Date()}/>
}