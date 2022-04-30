import {createSlice} from "@reduxjs/toolkit";

export interface ScheduleState {
}

const initialScheduleState: ScheduleState = {
}

export const scheduleSlice = createSlice({
    name: "schedule",
    initialState: initialScheduleState,
    reducers: {
    }
})