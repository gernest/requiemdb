import { Box, Button, Heading, IconButton, PageLayout, Tooltip } from '@primer/react';
import { Editor } from '../../components';
import { FileIcon, SidebarCollapseIcon, SidebarExpandIcon } from '@primer/octicons-react';
import { useState } from 'react';


export const Console = () => {
    const [expand, setExpand] = useState<boolean>(false)
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Content>
                <Box display={"flex"}
                    justifyContent={"space-between"}
                    maxWidth={"100%"}
                    mb={3}
                >

                    <Box display={"flex"} alignItems={"center"}
                    >
                        <Box mr={1}>
                            {!expand && (
                                <Tooltip text="Expand file tree">
                                    <IconButton aria-label='Expand file tree' icon={SidebarExpandIcon}
                                        onClick={() => setExpand(true)} />
                                </Tooltip>
                            )}
                            {expand && (
                                <Tooltip text="Collapse file tree">
                                    <IconButton aria-label='Collapse file tree' icon={SidebarCollapseIcon}
                                        onClick={() => setExpand(false)} />
                                </Tooltip>
                            )}
                        </Box>
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
            </PageLayout.Content>
            {expand && (<PageLayout.Pane position={"start"} width={"small"}>
                Tree
            </PageLayout.Pane>)}
        </PageLayout>

    )
}