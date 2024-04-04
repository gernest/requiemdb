import { Box, Button, PageLayout, Text } from '@primer/react';
import { Editor } from '../../components';
import { TriangleRightIcon } from '@primer/octicons-react';


export const Console = () => {
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Header>
                <Box display={"flex"}
                    justifyContent={"space-between"}
                    maxWidth={"100%"}
                >
                    <Box
                        display={"flex"}
                        alignItems={"center"}
                        flexDirection={"row"}
                        px={1}
                    >
                        <Button leadingVisual={TriangleRightIcon}>
                            <Text color={"accent.emphasis"}>Run</Text>
                        </Button>
                    </Box>
                </Box>
            </PageLayout.Header>
            <PageLayout.Content>
                <Editor width={"100%"} minHeight={400} />
            </PageLayout.Content>
        </PageLayout>

    )
}


