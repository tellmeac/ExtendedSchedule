import React, {ReactNode, useEffect, useMemo, useState} from "react";
import {ScheduleDay} from "../../Shared/Models";
import {Table} from "react-bootstrap";
import {format} from "date-fns";
import "./WeekSchedule.css"
import {Intervals, IntervalSectionsCount} from "../../Shared/Definitions";
import {LessonCell} from "../LessonCell";
import 'bootstrap/dist/css/bootstrap.min.css';
import {useAppDispatch, useAppSelector} from "../../Shared/Hooks";
import {selectWeekPeriod, selectWeekSchedule, updateSchedule} from "../../Shared/Store";
import {selectLoginResponse} from "../../Shared/Store";
import {getUserSchedule} from "../../Shared/Api";

interface ColumnInfo {
    weekDay: string
    shortDate: string
}

function generateColumnsInfo(days: ScheduleDay[]): ColumnInfo[] {
    return days.map((day) => {
        const date = new Date(day.date)
        return {
            weekDay: format(date, "E"),
            shortDate: format(date, "d LLL.")
        }
    });
}

export const WeekSchedule: React.FC = () => {
    const dispatch = useAppDispatch()

    const weekPeriod = useAppSelector(selectWeekPeriod)
    const weekSchedule = useAppSelector(selectWeekSchedule)
    const user = useAppSelector(selectLoginResponse)

    const [columnWeekDayInfo, setColumnWeekDayInfo] = useState<ColumnInfo[]>([]);

    const getWeekSchedule = () => {
        if (!user) {
            return
        }

        getUserSchedule(user?.tokenId || "", weekPeriod.weekStart, weekPeriod.weekEnd).then(days => {
            dispatch(updateSchedule(days))
        }).catch(e => {
            console.log(e)
        })
    }

    useEffect(() => {
        getWeekSchedule()
    }, [weekPeriod, user])

    useMemo(() => {
        setColumnWeekDayInfo(generateColumnsInfo(weekSchedule))
    }, [weekSchedule])

    return <Table striped bordered hover>
        <thead>
        <tr>
            <th key="-1"/>
            {
                columnWeekDayInfo.map((colInfo)=>{
                    return <th key={colInfo.weekDay}>
                        <div className={"day-header"}>
                            <span className={"day-header-weekday"}>{colInfo.weekDay}</span>
                            <span className={"day-header-date"}>{colInfo.shortDate}</span>
                        </div>
                    </th>
                })
            }
        </tr>
        </thead>
        <tbody>
        {
            Array.from(Array<boolean>(IntervalSectionsCount).keys()).map(position => {
                return <tr key={position}>
                    <td key={-1}>
                        <div className={"section-header"}>
                            <span key="start-date" className={"section-start-date"}>{Intervals[position][0]}</span>
                            <span key="end-date" className={"section-end-date"}>{Intervals[position][1]}</span>
                        </div>
                    </td>
                    {renderSection(position, weekSchedule)}
                </tr>
            })
        }
        </tbody>
    </Table>
}

function renderSection(position: number, days: ScheduleDay[]): ReactNode {
    return <>
        {
            days.map(day => {
                return <td key={day.date}>
                    {
                        day.lessons.filter(lesson => {
                            return lesson.position === position
                        }).map(lesson => {
                            return <LessonCell key={lesson.position + lesson.id}
                                               lesson={lesson}/>
                        })
                    }
                </td>
            })
        }
    </>
}