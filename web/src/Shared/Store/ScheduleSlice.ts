import {createSlice} from "@reduxjs/toolkit";
import {RootState} from "./Root";

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

// export const selectSubjects = (state: RootState) => state.schedule.subjects
// export const selectChosenArticle = (state: RootState) => state.schedule.chosenArticle
// export const selectNewSubjectData = (state: RootState) => state.schedule.newSubject

// export const { updateSubjectsList, setChosenArticle, setNewSubjectData } = scheduleSlice.actions