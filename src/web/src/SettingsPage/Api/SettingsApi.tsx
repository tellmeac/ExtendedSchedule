import {FacultyInfo, GroupInfo} from "../../Shared/Models";
import axios from "axios";
import {LessonInfo, UserConfig} from "../Models";

const base = process.env.REACT_APP_API_BASE_URL

/**
 * Returns user configuration
 */
export async function getUserConfig(): Promise<UserConfig> {
    return (await axios.get<UserConfig>(`${base}/user/config`)).data
}

/**
 * Updates user configuration
 */
export async function updateUserConfig(desired: UserConfig): Promise<void> {
    await axios.patch(`${base}/user/config`, desired)
}

/**
 * Returns nearest group lessons
 * @param groupId - group identifier
 */
export async function getLessonsInfo(groupId: string): Promise<LessonInfo[]> {
    return (
        await axios.get<LessonInfo[]>(`${base}/lessons/`, {
            params: {
                "groupId": groupId
            }
        })
    ).data
}

/**
 * Returns all faculties
 */
export async function getAllFaculties(): Promise<FacultyInfo[]> {
    return (
        await axios.get<FacultyInfo[]>(`${base}/faculties/`)
    ).data
}

/**
 * Returns groups of specific faculty
 * @param facultyId - faculty identifier
 */
export async function getFacultyGroups(facultyId: string): Promise<GroupInfo[]> {
    return (await axios.get<GroupInfo[]>(`${base}/faculties/${facultyId}/groups`)).data
}