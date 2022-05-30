import {GroupInfo} from "../../Shared/Models";

/**
 * User configuration
 */
export interface UserConfig {
    id: string
    email: string
    baseGroup: GroupInfo
    excludedLessons: ExcludedRule[]
    extendedGroupLessons: ExtendedLessons[]
}

/**
 * Exclude rule for lessons
 * TODO: feature
 */
export interface ExcludedRule {}

/**
 * ExtendedLessons is set of lessons to include from group to user's schedule
 */
export interface ExtendedLessons {
    group: GroupInfo
    lessonIds: string[]
}

/**
 * LessonInfo without schedule day context
 */
export interface LessonInfo {
    id: string
    name: string
    teacherName: string
    lessonKind: string
}