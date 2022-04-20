import React, {useMemo} from "react";
import {EmptyCell, ScheduleDay} from "../../Shared/Models";
import {Table} from "react-bootstrap";
import { format } from "date-fns";
import "./WeekScheduleTable.css"
import {Intervals, ScheduleWeekDayCount, SectionsCount} from "../../Shared/Constants";
import {MockScheduleWeek} from "../Mocks/MockScheduleData";
import {LessonCell} from "../Components/LessonCell";


interface WeekScheduleProps {
    dateStart: Date
    dateEnd: Date
    days: ScheduleDay[]
}

interface ColumnInfo {
    weekDay: string
    shortDate: string
}

function generateColumnsInfo(days: ScheduleDay[]): ColumnInfo[] {
    return days.map((day) => {
        return {
            weekDay: format(day.date, "E"),
            shortDate: format(day.date, "d LLL.")
        }
    });
}

export const WeekScheduleTable: React.FC<WeekScheduleProps> = ({dateStart, dateEnd, days}) => {
    const columnsInfo: ColumnInfo[] = generateColumnsInfo(days)

    const week = useMemo(() => MockScheduleWeek, [])
    console.log(week)

    return <Table striped bordered hover>
        <thead>
            <tr>
                <th> # </th>
                {
                    columnsInfo.map((colInfo)=>{
                        return <th key={colInfo.weekDay}>
                            <p>{colInfo.weekDay}</p>
                            <p>{colInfo.shortDate}</p>
                        </th>
                    })
                }
            </tr>
        </thead>
        <tbody>
            {
                // rendering table body section by section
                Array.from(Array<boolean>(SectionsCount).keys()).map((sectionNumber) => {
                    return <tr key={sectionNumber}>
                        <td>
                            <p className={"date-start-section"}>{Intervals[sectionNumber][0]}</p>
                            <p>{Intervals[sectionNumber][1]}</p>
                        </td>
                        {
                            Array.from(Array<boolean>(ScheduleWeekDayCount).keys()).map((weekDay) => {
                                // console.log(`week = ${weekDay}, section = ${sectionNumber}`)
                                return <td key={weekDay}>
                                    {
                                        week[weekDay].sections[sectionNumber].lessons.map((lesson) => {
                                            return <>
                                                {
                                                    lesson.type !== EmptyCell &&
                                                    <LessonCell lesson={lesson}/>
                                                }
                                            </>
                                        })
                                    }
                                </td>
                            })
                        }
                    </tr>
                })
            }
        </tbody>
    </Table>
}