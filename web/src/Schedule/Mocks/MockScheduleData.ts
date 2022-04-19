import {DaySection, Lesson, ScheduleDay} from "../../Shared/Models";
import {add, startOfWeek} from "date-fns";

export const MockScheduleWeek: ScheduleDay[] =  generateCurrentWeek()

function generateCurrentWeek(): ScheduleDay[] {
    const result: ScheduleDay[] = []
    const curr = new Date()
    const mondayDate = startOfWeek(curr, {weekStartsOn: 1})

    for(let i = 0; i < 7; i++){
        result.push(
            {
                date: add(mondayDate, {days: i}),
                sections: generateSections()
            }
        )
    }
    return result
}


function generateSections(): DaySection[] {
    const result: DaySection[] = []

    const choices = ["practice", "lesson", "seminar"]

    for(let i = 0; i < 8; i++){
        if (Math.random() <= 0.55) {
            result.push(
                {
                    position: i,
                    lessons: [
                        {
                            id: i.toString(),
                            title: "Study",
                            type: choose(choices),
                            audience: {
                                id: "zxc",
                                name: "online"
                            },
                            groups: [
                                {
                                    id: "zxc",
                                    name: "931901"
                                }
                            ],
                            professor: {
                                id: "zxc",
                                name: "Unknown"
                            }
                        }
                    ]
                }
            )
        } else {
            result.push(
                {
                    position: i,
                    lessons: [
                        {
                            id: i.toString(),
                            title: "Empty",
                            type: "empty"
                        }
                    ]
                }
            )
        }
    }

    return result
}

function choose(choices: any[]): any {
    return choices[Math.floor(Math.random() * choices.length)];
}