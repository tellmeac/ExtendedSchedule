import {configureStore} from '@reduxjs/toolkit';
import {scheduleSlice} from "./ScheduleSlice";
import {uiSlice} from "./UISlice";
import {userSlice} from "./UserSlice";

export const store = configureStore({
    reducer: {
        schedule: scheduleSlice.reducer,
        ui: uiSlice.reducer,
        user: userSlice.reducer,
    },
})

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch