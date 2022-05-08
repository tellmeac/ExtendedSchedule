import {ScheduleDay} from "../Models";
import {format} from "date-fns";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

export async function getUserSchedule(tokenId: string, startTime: number, endTime: number): Promise<ScheduleDay[]> {
    const start = format(new Date(startTime), "u-MM-dd")
    const end = format(new Date(endTime), "u-MM-dd")

    const url = `${ScheduleAPIBaseUrl}/schedule/personal?start=${start}&end=${end}`
    const response = fetch(url, {
        method: 'GET',
        headers: {
            "Authorization": `Bearer ${tokenId}`,
        },
    });

    return response.then(r => {
        return r.json() as Promise<ScheduleDay[]>
    }).catch(err => {
        throw err
    })
}