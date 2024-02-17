import { AppBar, CssBaseline, Drawer, Toolbar, Typography } from "@mui/material";
import { Box } from "@mui/system";
import React from "react";
import { Outlet } from "react-router";

const Layout = (): JSX.Element => (
    <Box sx={{ display: "flex" }}>
        <CssBaseline />
        <AppBar position="fixed" sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}>
            <Toolbar>
                <Typography variant="h5" component="h1">
                    UWR Tournaments
                </Typography>
            </Toolbar>
        </AppBar>
        <Drawer
            variant="permanent"
            sx={{
                width: 240,
                flexShrink: 0,
                [`& .MuiDrawer-paper`]: { width: 240, boxSizing: "border-box" },
            }}
        >
            <Toolbar />
        </Drawer>

        <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
            <Toolbar />
            <Outlet />
        </Box>
    </Box>
);

export default Layout;
