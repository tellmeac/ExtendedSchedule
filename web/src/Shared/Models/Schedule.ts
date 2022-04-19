export const EmptyCell: string = "empty"
export const LectureCell: string = "lesson"
export const PracticeCell: string = "practice"
export const SeminarCell: string = "seminar"

export interface ScheduleDay {
    date: Date
    sections: DaySection[]
}

export interface DaySection {
    position: number
    lessons: Lesson[]
}

export interface Lesson {
    id: string
    title: string
    type: string
    audience?: Audience
    groups?: Group[]
    professor?: Professor
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

