import {add, startOfWeek} from "date-fns";

/**
 * Returns current week monday regarding to passed day
 * @param current
 */
export function getCurrentWeekMonday(current: Date): Date {
    return startOfWeek(current, {weekStartsOn: 1})
}

/**
 * Returns current week saturday regarding to passed day
 * @param current
 */
export function getCurrentWeekSaturday(current: Date): Date {
    return add(getCurrentWeekMonday(current), {days: 5})
}