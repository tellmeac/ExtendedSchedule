import {ScheduleDay} from "../../Shared/Models";
import {format} from "date-fns";
import axios from "axios";

const base = process.env.REACT_APP_API_BASE_URL

/**
 * Receives user's personal schedule
 * @param startTime start day of schedule
 * @param endTime end day of schedule
 */
export async function getPersonalSchedule(startTime: number, endTime: number): Promise<ScheduleDay[]> {
    return (
        await axios.get<ScheduleDay[]>(`${base}/schedule/personal`, {
            params: {
                "start": format(new Date(startTime), "u-MM-dd"),
                "end": format(new Date(endTime), "u-MM-dd"),
            },
        })
    ).data
}