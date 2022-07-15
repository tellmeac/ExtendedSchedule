import React from "react";
import {ScheduleDay} from "../../../Shared/Models";
import {Table} from "react-bootstrap";
import {add, format} from "date-fns";
import "./WeekSchedule.css"
import {Intervals, IntervalSectionsCount} from "../../../Shared/Definitions";
import {LessonCell} from "../LessonCell";
import 'bootstrap/dist/css/bootstrap.min.css';

type WeekScheduleProps = {
    startDay: number
    endDay: number
    schedule: ScheduleDay[]
}

/**
 * WeekSchedule is a component to render passed week schedule
 * @constructor
 */
export const WeekSchedule: React.FC<WeekScheduleProps> = ({startDay, endDay, schedule}) => {
    return <Table striped bordered hover>
        <thead>
        <tr>
            <th key="-1"/>
            {
                dayInfo(startDay, endDay).map((colInfo)=>{
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
                    <ScheduleSection pos={position} days={schedule}/>
                </tr>
            })
        }
        </tbody>
    </Table>
}

/**
 * Content for column header in a week schedule table
 */
interface ColumnInfo {
    weekDay: string
    shortDate: string
}

/**
 * Generates day header info for table header
 * @param startDay
 * @param endDay
 */
function dayInfo(startDay: number, endDay: number): ColumnInfo[] {
    const dates: number[] = []
    let end = new Date(endDay)
    let d = new Date(startDay)
    while (d <= end) {
        dates.push(d.getTime())
        d = add(d, {days: 1})
    }

    return dates.map((d) => {
        const date = new Date(d)
        return {
            weekDay: format(date, "E"),
            shortDate: format(date, "d LLL.")
        }
    });
}

/**
 * Renders schedule section with passed whole days by position
 * @param pos is a section number
 * @param days is a whole days, that will be rendered as table row by required position
 * @constructor
 */
const ScheduleSection: React.FC<{pos: number, days: ScheduleDay[]}> = ({pos, days}) => {
    return <>
        {
            days.map(day => {
                return <td key={day.date}>
                    {
                        day.lessons.filter(lesson => {
                            return lesson.position === pos
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