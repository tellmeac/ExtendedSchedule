import {laboratoryCell, LectureCell, PracticeCell, SeminarCell} from "../../../Shared/Models";
import React from "react";
import {LessonInfo} from "../../Models";
import "./LessonInfo.css"

interface Props {
    lesson: LessonInfo
}

export const LessonItem: React.FC<Props> = ({lesson}) => {
    let specificCellClass = ""
    switch (lesson.lessonKind) {
        case PracticeCell:
            specificCellClass = "practice-cell"
            break
        case LectureCell:
            specificCellClass = "lecture-cell"
            break
        case SeminarCell:
            specificCellClass = "seminar-cell"
            break
        case laboratoryCell:
            specificCellClass = "laboratory-cell"
            break
    }

    return <div className={["cell", specificCellClass].join(" ")}>
        <span className={"lesson-title"}>{lesson.name}</span>
        <span>{lesson.teacherName}</span>
        <span>{lesson.lessonKind.toLowerCase()}</span>
    </div>
}