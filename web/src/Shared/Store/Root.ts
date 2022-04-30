import { configureStore } from '@reduxjs/toolkit';
import {scheduleSlice} from "./ScheduleSlice";
import {uiSlice} from "./UISlice";

export const store = configureStore({
    reducer: {
        schedule: scheduleSlice.reducer,
        ui: uiSlice.reducer,
    },
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch