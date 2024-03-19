import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Layout from "./layout/Layout";
import Login from "./components/Login";

const Application = (): JSX.Element => (
    <BrowserRouter>
        <Routes>
            <Route path="/v2" element={<Layout />}>
                <Route index element={<Login />} />
            </Route>
        </Routes>
    </BrowserRouter>
);

export default Application;
