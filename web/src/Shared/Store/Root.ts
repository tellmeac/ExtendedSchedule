import { configureStore } from '@reduxjs/toolkit';
import {scheduleSlice} from "./ScheduleSlice";

export const store = configureStore({
    reducer: {
        schedule: scheduleSlice.reducer
    },
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch