import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Layout from "./layout/Layout";
import UserShow from "./user/UserShow";

const Application = (): JSX.Element => (
    <BrowserRouter>
        <Routes>
            <Route path="/v2" element={<Layout />}>
                <Route index element={<UserShow />} />
                <Route path="tournament" element={<div>This would be in a tournament</div>} />
            </Route>
        </Routes>
    </BrowserRouter>
);

export default Application;
