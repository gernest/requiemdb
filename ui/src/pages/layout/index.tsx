import { Outlet } from "react-router-dom";

import { Footer, Sidebar, MainHeader } from "../../components";
import { Box, PageLayout } from '@primer/react';

export const Layout = () => {
    return (
        <Box
            minHeight={"100vh"}
            overflow={"auto"}
        >
            <MainHeader />
            <PageLayout containerWidth="full">
                <PageLayout.Content>
                    <Outlet />
                </PageLayout.Content>
                <PageLayout.Footer>
                    <Footer />
                </PageLayout.Footer>
            </PageLayout>
        </Box>
    )
}

