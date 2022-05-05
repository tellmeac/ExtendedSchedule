import {LectureCell, Lesson, PracticeCell, ScheduleDay, SeminarCell} from "../../Shared/Models";
import {add} from "date-fns";
import {LessonCell} from "../LessonCell";

export function generateWeekSchedule(monday: number): ScheduleDay[] {
    const result: ScheduleDay[] = []

    for(let i = 0; i < 6; i++){
        result.push(
            {
                date: add(new Date(monday), {days: i}).getTime(),
                lessons: generateLessons()
            }
        )
    }
    return result
}


function generateLessons(): Lesson[] {
    const result: Lesson[] = []

    const choices = [PracticeCell, LectureCell, SeminarCell]

    function generateGroups() {
        const result = [{
            id: "g1",
            name: "931901"
        }]
        if (Math.random() <= 0.3) {
            result.push({
                id: "g2",
                name: "931901 a"
            })
        }
        return result;
    }

    function generateLesson(id: string, position: number): Lesson {
        return {
            id: id,
            title: "Study Subject",
            position: position,
            lessonType: choose(choices),
            audience: {
                id: "zxc",
                name: "online"
            },
            groups: generateGroups(),
            professor: {
                id: "zxc",
                name: "Unknown"
            }
        };
    }

    const randomId = () => {
        return (Math.random() + 1).toString(36)
    }

    for(let i = 0; i < 8; i++){
        if (Math.random() <= 0.45) {
            result.push(generateLesson(randomId(), i))

            if (Math.random() <= 0.3) {
                result.push(generateLesson(randomId(), i))
            }
        }
    }

    return result
}

function choose(choices: any[]): any {
    return choices[Math.floor(Math.random() * choices.length)];
}