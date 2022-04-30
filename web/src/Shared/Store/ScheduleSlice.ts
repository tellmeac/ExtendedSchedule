import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {ScheduleDay} from "../Models";
import {getCurrentWeekMonday, getCurrentWeekSaturday} from "../Definitions";
import {add, sub} from "date-fns";
import {RootState} from "./Root";

export interface ScheduleState {
    selectedWeekStart: Date
    selectedWeekEnd: Date
    weekSchedule: ScheduleDay[]
}

const initialScheduleState: ScheduleState = {
    selectedWeekStart: getCurrentWeekMonday(new Date()),
    selectedWeekEnd: getCurrentWeekSaturday(new Date()),
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
            state.selectedWeekStart = add(state.selectedWeekStart, {days: 7})
            state.selectedWeekEnd = add(state.selectedWeekEnd, {days: 7})
        },
        setPreviousWeek: (state) => {
            state.selectedWeekStart = sub(state.selectedWeekStart, {days: 7})
            state.selectedWeekEnd = sub(state.selectedWeekEnd, {days: 7})
        },
    }
})

export const selectSelectedWeekStart = (state: RootState) => state.schedule.selectedWeekStart
export const selectSelectedWeekEnd = (state: RootState) => state.schedule.selectedWeekEnd
export const selectSelectedSchedule = (state: RootState) => state.schedule.weekSchedule

export const { updateSchedule, setNextWeek, setPreviousWeek, } = scheduleSlice.actions