import { Box, Button, PageLayout } from '@primer/react';
import { Editor } from '../../components';
import { FileIcon } from '@primer/octicons-react';


export const Console = () => {
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Header>
                <Box display={"flex"}
                    justifyContent={"space-between"}
                    maxWidth={"100%"}
                >
                    <Box display={"flex"} alignItems={"center"}
                    >
                        <Button
                            leadingVisual={FileIcon}
                        >http_request_total.ts</Button>

                    </Box>
                    <Box
                        display={"flex"}
                        alignItems={"center"}
                        flexDirection={"row"}
                        px={1}
                    >
                        <Button>cancel changes</Button>
                        <Button sx={{ marginLeft: "12px" }} variant='primary'>save changes</Button>
                    </Box>
                </Box>
            </PageLayout.Header>
            <PageLayout.Content>
                <Editor width={"100%"} minHeight={400} />
            </PageLayout.Content>
            <PageLayout.Pane position={"start"}>
                Tree
            </PageLayout.Pane>
        </PageLayout>

    )
}