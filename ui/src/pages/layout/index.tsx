import { Outlet } from "react-router-dom";

import { Footer, Sidebar, MainHeader } from "../../components";
import { Box, PageLayout } from '@primer/react';

export const Layout = () => {
    return (
        <Box
            minHeight={"100vh"}
            display={"grid"}
            gridTemplateRows={"65px calc(100vh - 110px) 45px"}
            overflow={"hidden"}
        >
            <MainHeader />
            <PageLayout containerWidth='full'>
                <PageLayout.Pane position={"start"} sticky
                    width={{
                        min: "126px",
                        default: "126px",
                        max: "126px",
                    }}
                    divider={"line"}
                >
                    <Sidebar />
                </PageLayout.Pane>
                <PageLayout.Content>
                    <Outlet />
                </PageLayout.Content>
            </PageLayout>
            <Footer />
        </Box>
    )
}