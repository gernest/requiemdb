import { Box, Button } from '@primer/react';
import { Editor } from '../../components';
import { FileIcon } from '@primer/octicons-react';


export const Console = () => {
    return (
        <Box>
            <Box display={"flex"}
                justifyContent={"space-between"}
                mb={3}
                pb={2}
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
            <Editor width={"100%"} minHeight={400} />
        </Box>
    )
}