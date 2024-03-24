import styled from 'styled-components';
import { Footer, Sidebar, Editor } from "../../components";
import { PageLayout } from '@primer/react';
const Wrapper = styled.div``


export const Layout = () => {
    return (
        <Wrapper>
            <PageLayout containerWidth='full'>
                <PageLayout.Header></PageLayout.Header>
                <PageLayout.Pane position={"start"} resizable>
                    <Sidebar />
                </PageLayout.Pane>
                <PageLayout.Content>
                    <Editor height={500} />
                </PageLayout.Content>
                <PageLayout.Footer>
                    <Footer />
                </PageLayout.Footer>
            </PageLayout>
        </Wrapper>
    )
}