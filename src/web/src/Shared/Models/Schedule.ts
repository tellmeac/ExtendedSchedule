export const LectureCell: string = "LECTURE"
export const PracticeCell: string = "PRACTICE"
export const SeminarCell: string = "SEMINAR"
export const laboratoryCell: string = "LABORATORY"

/**
 * Schedule day representation
 */
export interface ScheduleDay {
    date: number
    lessons: Lesson[]
}

/**
 * Lesson model
 */
export interface Lesson {
    id: string
    title: string
    position: number
    lessonType: string
    audience: Audience
    groups: GroupInfo[]
    professor: Professor
}

/**
 * Audience info
 */
export interface Audience {
    id: string
    name: string
}

/**
 * Group Info
 */
export interface GroupInfo {
    id: string
    name: string
}


/**
 * Faculty Info
 */
export interface FacultyInfo {
    id: string
    name: string
}

/**
 * Professor info
 */
export interface Professor {
    id: string
    name: string
}

