import {FacultyInfo} from "../../Shared/Models";
import axios from "axios";
import {applyAuthorization} from "../../Shared/Api/Token";

const ScheduleAPIBaseUrl = process.env.REACT_APP_API_BASE_URL || "http://localhost:8080/api"

/**
 * Returns all faculties
 */
export async function getAllFaculties(): Promise<FacultyInfo[]> {
    const config = applyAuthorization({
        validateStatus: status => {
            return status < 400
        }
    })

    const response = await axios.get<FacultyInfo[]>(`${ScheduleAPIBaseUrl}/faculties/`,config)

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