import {ScheduleDay} from "../../Shared/Models";
import {format} from "date-fns";
import axios from "axios";
import {applyAuthorization} from "../../Shared/Auth/Token";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

/**
 * Receives user's personal schedule
 * @param startTime start day of schedule
 * @param endTime end day of schedule
 */
export async function getPersonalSchedule(startTime: number, endTime: number): Promise<ScheduleDay[]> {
    const start = format(new Date(startTime), "u-MM-dd")
    const end = format(new Date(endTime), "u-MM-dd")

    const config = applyAuthorization({
        params: {
            "start": start,
            "end": end
        },
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<ScheduleDay[]>(`${ScheduleAPIBaseUrl}/schedule/personal`, config)
    return response.data
}