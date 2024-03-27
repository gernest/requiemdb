import { Outlet } from "react-router-dom";

import { Footer, Sidebar, MainHeader } from "../../components";
import { Box } from '@primer/react';

export const Layout = () => {
    return (
        <Box
            minHeight={"100vh"}
            display={"grid"}
            gridTemplateRows={"65px calc(100vh - 110px) 45px"}
            overflow={"hidden"}
        >
            <MainHeader />
            <Box height={"100%"}
                display={"grid"}
                gridTemplateColumns={"80px calc(100vw - 80px)"}
                overflow={"hidden"}
            >
                <Sidebar />
                <Box width={"100%"} height={"100%"}>
                    <Outlet />
                </Box>
            </Box>
            <Footer />
        </Box>
    )
}

