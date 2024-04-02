import { Outlet } from "react-router-dom";

import { Footer, MainHeader } from "../../components";
import { Box } from '@primer/react';

export const Layout = () => {
    return (
        <Box>
            <MainHeader />
            <Outlet />
            <Footer />
        </Box>
    )
}

