import React, {ReactNode, useEffect, useMemo, useState} from "react";
import {EmptyCell, ScheduleDay} from "../../Shared/Models";
import {Table} from "react-bootstrap";
import {format} from "date-fns";
import "./WeekScheduleTable.css"
import {Intervals, IntervalSectionsCount} from "../../Shared/Constants";
import {generateCurrentWeek} from "../Mocks/MockScheduleData";
import {LessonCell} from "../Components";
import 'bootstrap/dist/css/bootstrap.min.css';


interface WeekScheduleProps {
    dateStart: Date
    dateEnd: Date
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

export const WeekScheduleTable: React.FC<WeekScheduleProps> = ({dateStart, dateEnd}) => {
    let columnsInfo: ColumnInfo[] = [];
    const [scheduleDays, setScheduleDays] = useState<ScheduleDay[]>([]);

    const updateSchedule = () => {
        setScheduleDays(generateCurrentWeek())
    }

    useEffect(() => {
        updateSchedule()
    }, [])

    useMemo(() => {
        columnsInfo = generateColumnsInfo(scheduleDays)
    }, [scheduleDays])

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
            Array.from(Array<boolean>(IntervalSectionsCount).keys()).map(position => {
                return <tr key={position}>
                    <td key={-1}>
                        <p key="start-date" className={"date-start-section"}>{Intervals[position][0]}</p>
                        <p key="end-date">{Intervals[position][1]}</p>
                    </td>
                    {renderSection(position, scheduleDays)}
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
                return <td key={day.date.toString()}>
                    {
                        day.lessons.filter(lesson => {
                            return lesson.position === position
                        }).map(lesson => {
                            const key = lesson.id + lesson.position + lesson.professor.id
                            return <div key={key}>
                                {
                                    lesson.type !== EmptyCell &&
                                    <LessonCell key={key}
                                                lesson={lesson}/>
                                }
                            </div>
                        })
                    }
                </td>
            })
        }
    </>
}