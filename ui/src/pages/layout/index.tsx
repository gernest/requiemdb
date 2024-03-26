import { Outlet } from "react-router-dom";

import { Footer, Sidebar, MainHeader } from "../../components";
import { PageLayout } from '@primer/react';

export const Layout = () => {
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Header>
                <MainHeader />
            </PageLayout.Header>
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
            <PageLayout.Footer>
                <Footer />
            </PageLayout.Footer>
        </PageLayout>
    )
}