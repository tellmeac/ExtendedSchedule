import {createSlice} from "@reduxjs/toolkit";

export interface UIState {
}

const initialUIState: UIState = {
}

export const uiSlice = createSlice({
    name: "ui",
    initialState: initialUIState,
    reducers: {
    }
})