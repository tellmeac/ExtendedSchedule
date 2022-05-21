import {FacultyInfo, ScheduleDay} from "../../Shared/Models";
import axios from "axios";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

/**
 * Returns all faculties
 * @param tokenId
 */
export async function getAllFaculties(tokenId: string): Promise<FacultyInfo[]> {
    const response = await axios.get<FacultyInfo[]>(`${ScheduleAPIBaseUrl}/faculties/`,{
        headers: {
            "Authorization": `Bearer ${tokenId}`,
        },
        validateStatus: status => {
            return status < 400
        }
    })

    return response.data
}

/**
 * Returns groups of specific faculty
 * @param tokenId
 * @param facultyId
 */
export async function getFacultyGroups(tokenId: string, facultyId: string): Promise<FacultyInfo[]> {
    const response = await axios.get<FacultyInfo[]>(`${ScheduleAPIBaseUrl}/faculties/${facultyId}/groups`,{
        headers: {
            "Authorization": `Bearer ${tokenId}`,
        },
        validateStatus: status => {
            return status < 400
        }
    })

    return response.data
}