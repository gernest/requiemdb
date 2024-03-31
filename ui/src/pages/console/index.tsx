import { Box, Button, Heading, IconButton, PageLayout, Text, Tooltip } from '@primer/react';
import { Editor } from '../../components';
import { FileIcon, PlusIcon, SidebarCollapseIcon, SidebarExpandIcon, TriangleRightIcon } from '@primer/octicons-react';
import { useState } from 'react';


export const Console = () => {
    const [expand, setExpand] = useState<boolean>(false)
    return (
        <PageLayout containerWidth='full'>
            <PageLayout.Header>
                <Box display={"flex"}
                    justifyContent={"space-between"}
                    maxWidth={"100%"}
                >

                    <Box display={"flex"} alignItems={"center"}
                    >
                        <Box mr={1}>
                            {!expand && (
                                <Tooltip text="Expand file tree" direction='se'>
                                    <IconButton aria-label='Expand file tree' icon={SidebarExpandIcon}
                                        onClick={() => setExpand(true)} />
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
                        <Button leadingVisual={TriangleRightIcon}>
                            <Text color={"accent.emphasis"}>Run</Text>
                        </Button>
                        <Button sx={{ marginLeft: "12px" }} variant='primary'>save changes</Button>
                    </Box>
                </Box>
            </PageLayout.Header>
            <PageLayout.Content>
                <Editor width={"100%"} minHeight={400} />
            </PageLayout.Content>
            {expand && (<PageLayout.Pane position={"start"} width={"small"} sx={{ backgroundColor: "canvas.subtle" }}>
                <Box display={"flex"} flexDirection={"column"} p={1}>
                    <Box display={"flex"} alignItems={"center"} flexDirection={"column"}>
                        <Box display={"flex"} alignItems={"center"} width={"100%"}>
                            <Heading sx={{ fontSize: "14px", display: "flex", margin: 0, fontWeight: 600 }}>
                                <Tooltip text="Collapse file tree" direction='se'>
                                    <IconButton aria-label='Collapse file tree' icon={SidebarCollapseIcon}
                                        onClick={() => setExpand(false)} />
                                </Tooltip>

                            </Heading>

                            <Heading sx={{
                                fontSize: "16px",
                                margin: "0px 0px 0px 8px",
                                display: "flex", fontWeight: 600,
                                flexGrow: 1,
                            }}>
                                Files
                            </Heading>

                            <Box>
                                <Tooltip direction="w" text="Add new file">
                                    <IconButton icon={PlusIcon} aria-label='Add file' />
                                </Tooltip>
                            </Box>
                        </Box>


                    </Box>
                </Box>
            </PageLayout.Pane>)}
        </PageLayout>

    )
}