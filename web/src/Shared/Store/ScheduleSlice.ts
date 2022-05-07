import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {ScheduleDay} from "../Models";
import {getCurrentWeekMonday, getCurrentWeekSaturday} from "../Definitions";
import {add, sub} from "date-fns";
import {RootState} from "./Root";

export interface ScheduleState {
    period: {
        weekStart: number,
        weekEnd: number
    }
    weekSchedule: ScheduleDay[]
}

const initialScheduleState: ScheduleState = {
    period: {
        weekStart: getCurrentWeekMonday(new Date()).getTime(),
        weekEnd: getCurrentWeekSaturday(new Date()).getTime()
    },
    weekSchedule: [],
}

export const scheduleSlice = createSlice({
    name: "schedule",
    initialState: initialScheduleState,
    reducers: {
        updateSchedule: (state, action: PayloadAction<ScheduleDay[]>) => {
            state.weekSchedule = action.payload
        },
        setNextWeek: (state) => {
            state.period = {
                weekStart: add(state.period.weekStart, {days: 7}).getTime(),
                weekEnd: add(state.period.weekEnd, {days: 7}).getTime()
            }
        },
        setPreviousWeek: (state) => {
            state.period = {
                weekStart: sub(state.period.weekStart, {days: 7}).getTime(),
                weekEnd: sub(state.period.weekEnd, {days: 7}).getTime()
            }
        },
    }
})

export const selectWeekPeriod = (state: RootState) => state.schedule.period
export const selectWeekSchedule = (state: RootState) => state.schedule.weekSchedule

export const { updateSchedule, setNextWeek, setPreviousWeek, } = scheduleSlice.actions