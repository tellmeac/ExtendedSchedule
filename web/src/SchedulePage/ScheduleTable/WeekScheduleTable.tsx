import React, {useMemo} from "react";
import {EmptyCell, ScheduleDay} from "../../Shared/Models";
import {Table} from "react-bootstrap";
import { format } from "date-fns";
import "./WeekScheduleTable.css"
import {Intervals, ScheduleWeekDayCount, SectionsCount} from "../../Shared/Constants";
import {MockScheduleWeek} from "../Mocks/MockScheduleData";
import {LessonCell} from "../Components";
import 'bootstrap/dist/css/bootstrap.min.css';


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

    return <Table striped bordered hover>
        <thead>
        <tr>
            <th key="-1"> # </th>
            {
                columnsInfo.map((colInfo)=>{
                    return <th key={colInfo.weekDay}>
                        <p>{colInfo.weekDay} {colInfo.shortDate}</p>
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
                    <td key="-1">
                        <p key="start-date" className={"date-start-section"}>{Intervals[sectionNumber][0]}</p>
                        <p key="end-date">{Intervals[sectionNumber][1]}</p>
                    </td>
                    {
                        Array.from(Array<boolean>(ScheduleWeekDayCount).keys()).map((weekDay) => {
                            // console.log(`week = ${weekDay}, section = ${sectionNumber}`)
                            return <td key={weekDay}>
                                {
                                    week[weekDay].sections[sectionNumber].lessons.map((lesson) => {
                                        const unique_key = [lesson.id, lesson.groups?.concat()].join("")
                                        return <div key={unique_key}>
                                            {
                                                lesson.type !== EmptyCell &&
                                                <LessonCell key={unique_key}
                                                            lesson={lesson}/>
                                            }
                                        </div>
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