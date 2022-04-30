import {Lesson, ScheduleDay} from "../../Shared/Models";
import {add, startOfWeek} from "date-fns";

export function generateCurrentWeek(): ScheduleDay[] {
    const result: ScheduleDay[] = []
    const curr = new Date()
    const mondayDate = startOfWeek(curr, {weekStartsOn: 1})

    for(let i = 0; i < 6; i++){
        result.push(
            {
                date: add(mondayDate, {days: i}),
                lessons: generateLessons()
            }
        )
    }
    return result
}


function generateLessons(): Lesson[] {
    const result: Lesson[] = []

    const choices = ["practice", "lesson", "seminar"]

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
            type: choose(choices),
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