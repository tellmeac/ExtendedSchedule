import React, {useEffect, useState} from "react";
import {WeekSchedule} from "./WeekSchedule";
import {ScheduleControlTab} from "./ScheduleControlTab/ScheduleControlTab";
import {Container} from "react-bootstrap";
import "./SchedulePage.css"
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {
    selectUserData,
    selectWeekPeriod,
    selectWeekSchedule,
    setNextWeek,
    setPreviousWeek,
    updateSchedule
} from "../Shared/Store";
import {getPersonalSchedule} from "../Shared/Api";

export function SchedulePage() {
    const dispatch = useAppDispatch()

    const user = useAppSelector(selectUserData)
    const period = useAppSelector(selectWeekPeriod)
    const schedule = useAppSelector(selectWeekSchedule)

    useEffect(()=>{
        getPersonalSchedule(user?.tokenId || "", period.weekStart, period.weekEnd).then((r)=>{
            dispatch(updateSchedule(r))
        }).catch((err)=>{
            console.error(err)
        })
    }, [period])

    const moveWeek = (isForward: boolean) => {
        if (isForward) {
            dispatch(setNextWeek())
            return
        }
        dispatch(setPreviousWeek())
    }

    return <Container className={"schedule-page"}>
        <ScheduleControlTab startDay={period.weekStart} endDay={period.weekEnd} moveWeek={moveWeek}/>
        <WeekSchedule startDay={period.weekStart} endDay={period.weekEnd} schedule={schedule}/>
    </Container>
}