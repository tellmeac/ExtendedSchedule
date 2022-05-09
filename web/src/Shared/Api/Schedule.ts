import {ScheduleDay} from "../Models";
import {format} from "date-fns";
import axios from "axios";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

/**
 * Receives user's personal schedule
 * @param tokenId raw JWT token
 * @param startTime start day of schedule
 * @param endTime end day of schedule
 */
export async function getPersonalSchedule(tokenId: string, startTime: number, endTime: number): Promise<ScheduleDay[]> {
    const start = format(new Date(startTime), "u-MM-dd")
    const end = format(new Date(endTime), "u-MM-dd")

    const response = await axios.get<ScheduleDay[]>(`${ScheduleAPIBaseUrl}/schedule/personal`,{
        params: {
          "start": start,
          "end": end
        },
        headers: {
            "Authorization": `Bearer ${tokenId}`,
        },
        validateStatus: status => {
            return status < 400
        }
    })
    return response.data
}