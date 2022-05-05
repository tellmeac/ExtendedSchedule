export const LectureCell: string = "LECTURE"
export const PracticeCell: string = "PRACTICE"
export const SeminarCell: string = "SEMINAR"
export const laboratoryCell: string = "LABORATORY"

export interface ScheduleDay {
    date: number
    lessons: Lesson[]
}

export interface Lesson {
    id: string
    title: string
    position: number
    lessonType: string
    audience: Audience
    groups: Group[]
    professor: Professor
}

export interface Audience {
    id: string
    name: string
}

export interface Group {
    id: string
    name: string
}

export interface Professor {
    id: string
    name: string
}

