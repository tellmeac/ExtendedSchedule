import React from "react";
import {laboratoryCell, LectureCell, Lesson, PracticeCell, SeminarCell} from "../../Shared/Models";
import "./LessonCell.css"

interface LessonCellProps {
    lesson: Lesson
}

export const LessonCell: React.FC<LessonCellProps> = ({lesson}) => {
    let specificCellClass = ""
    switch (lesson.lessonType) {
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
        <span className={"lesson-title"}>{lesson.title}</span>
        <span className={"lesson-group"}>{lesson.groups.map(g => g.name).join(", ")}</span>
    </div>
}