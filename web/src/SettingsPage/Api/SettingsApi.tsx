import {FacultyInfo} from "../../Shared/Models";
import axios from "axios";
import {applyAuthorization} from "../../Shared/Auth/Token";
import {LessonInfo, UserConfig} from "../Models";
import {format} from "date-fns";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

/**
 * Returns all faculties
 */
export async function getUserConfig(): Promise<UserConfig> {
    const config = applyAuthorization({
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<UserConfig>(`${ScheduleAPIBaseUrl}/user/config`, config)

    return response.data
}

/**
 * Returns nearest group lessons
 * @param groupId - group to get lessons from
 */
export async function getLessonsInfo(groupId: string): Promise<LessonInfo[]> {
    const config = applyAuthorization({
        params: {
            "groupId": groupId
        },
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<LessonInfo[]>(`${ScheduleAPIBaseUrl}/lessons/`, config)

    return response.data
}

/**
 * Returns all faculties
 */
export async function getAllFaculties(): Promise<FacultyInfo[]> {
    const config = applyAuthorization({
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<FacultyInfo[]>(`${ScheduleAPIBaseUrl}/faculties/`, config)

    return response.data
}

/**
 * Returns groups of specific faculty
 * @param facultyId
 */
export async function getFacultyGroups(facultyId: string): Promise<FacultyInfo[]> {
    const config = applyAuthorization({
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<FacultyInfo[]>(`${ScheduleAPIBaseUrl}/faculties/${facultyId}/groups`, config)
    return response.data
}