import {useEffect, useState} from "react";
import {WeekSchedule} from "./Components/WeekSchedule";
import {ScheduleControlTab} from "./Components/ScheduleControlTab/ScheduleControlTab";
import {Container} from "react-bootstrap";
import "./SchedulePage.css"
import {useAppDispatch, useAppSelector} from "../Shared/Hooks";
import {selectSignedIn, selectWeekPeriod, setNextWeek, setPreviousWeek} from "../Shared/Store";
import {getPersonalSchedule} from "./Api";
import {ScheduleDay} from "../Shared/Models";
import log from "loglevel";

export function SchedulePage() {
    const dispatch = useAppDispatch()

    const isAuthorized = useAppSelector(selectSignedIn)

    const period = useAppSelector(selectWeekPeriod)
    const [schedule, setSchedule] = useState<ScheduleDay[]>([])

    useEffect(()=>{
        if (!isAuthorized) {
            return
        }

        getPersonalSchedule(period.weekStart, period.weekEnd).then((r)=>{
            setSchedule(r)
        }).catch((err)=> {
            log.error(err)
        })
    }, [period, isAuthorized])

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