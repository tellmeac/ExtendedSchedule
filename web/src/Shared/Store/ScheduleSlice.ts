import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {ScheduleDay} from "../Models";
import {getCurrentWeekMonday, getCurrentWeekSaturday} from "../Definitions";
import {add, sub} from "date-fns";
import {RootState} from "./Root";

export interface ScheduleState {
    selectedWeekStart: number
    selectedWeekEnd: number
    weekSchedule: ScheduleDay[]
}

const initialScheduleState: ScheduleState = {
    selectedWeekStart: getCurrentWeekMonday(new Date()).getTime(),
    selectedWeekEnd: getCurrentWeekSaturday(new Date()).getTime(),
    weekSchedule: []
}

export const scheduleSlice = createSlice({
    name: "schedule",
    initialState: initialScheduleState,
    reducers: {
        updateSchedule: (state, action: PayloadAction<ScheduleDay[]>) => {
            state.weekSchedule = action.payload
        },
        setNextWeek: (state) => {
            state.selectedWeekStart = add(state.selectedWeekStart, {days: 7}).getTime()
            state.selectedWeekEnd = add(state.selectedWeekEnd, {days: 7}).getTime()
        },
        setPreviousWeek: (state) => {
            state.selectedWeekStart = sub(state.selectedWeekStart, {days: 7}).getTime()
            state.selectedWeekEnd = sub(state.selectedWeekEnd, {days: 7}).getTime()
        },
    }
})

export const selectSelectedWeekStart = (state: RootState) => state.schedule.selectedWeekStart
export const selectSelectedWeekEnd = (state: RootState) => state.schedule.selectedWeekEnd
export const selectSelectedSchedule = (state: RootState) => state.schedule.weekSchedule

export const { updateSchedule, setNextWeek, setPreviousWeek, } = scheduleSlice.actions