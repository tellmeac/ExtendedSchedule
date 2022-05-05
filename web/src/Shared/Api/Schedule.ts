import {ScheduleDay} from "../Models";
import {format} from "date-fns";

const ScheduleAPIBaseUrl = process.env.API_BASE_URL || "http://localhost:8080/api"

export async function getUserSchedule(id_token: string, startTime: number, endTime: number): Promise<ScheduleDay[]> {
    const start = format(new Date(startTime), "u-MM-dd")
    const end = format(new Date(endTime), "u-MM-dd")

    const url = `${ScheduleAPIBaseUrl}/schedule/groups/3c9f5a5d-ffca-11eb-8169-005056bc249c?start=${start}&end=${end}`
    const response = await fetch(url, {
        method: 'GET',
    });

    return (await response.json()) as ScheduleDay[]
}