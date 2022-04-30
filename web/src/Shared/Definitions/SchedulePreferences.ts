import {add, startOfWeek} from "date-fns";

export function getCurrentWeekMonday(current: Date): Date {
    return startOfWeek(current, {weekStartsOn: 1})
}

export function getCurrentWeekSaturday(current: Date): Date {
    return add(getCurrentWeekMonday(current), {days: 5})
}