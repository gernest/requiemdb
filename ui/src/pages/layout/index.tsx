import { Outlet } from "react-router-dom";

import { Footer, Sidebar, MainHeader } from "../../components";
import { PageLayout } from '@primer/react';

export const Layout = () => {
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Header>
                <MainHeader />
            </PageLayout.Header>
            <PageLayout.Pane position={"start"} resizable sticky>
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