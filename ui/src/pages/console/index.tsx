import { Box, Button, Dialog, FormControl, Heading, IconButton, PageLayout, Text, TextInput, Tooltip } from '@primer/react';
import { Editor } from '../../components';
import { FileIcon, PlusIcon, SidebarCollapseIcon, SidebarExpandIcon, TriangleRightIcon } from '@primer/octicons-react';
import { useRef, useState } from 'react';
import { SnippetInfo } from "../../rq";


export const Console = () => {
    const [expand, setExpand] = useState<boolean>(false)
    const [info, setInfo] = useState<SnippetInfo | undefined>()
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
                        <InfoBox info={info} />
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

                            <AddNewFile />
                        </Box>


                    </Box>
                </Box>
            </PageLayout.Pane>)}
        </PageLayout>

    )
}

const InfoBox = ({ info }: { info?: SnippetInfo }) => {
    return (
        <Box>
            <Button
                leadingVisual={FileIcon}
            >{info ? info.name : "(blank)"}</Button>
        </Box>
    )
}

const AddNewFile = () => {
    const [isOpen, setIsOpen] = useState(false)
    const focus = useRef(null)
    return (
        <Box>
            <Tooltip direction="w" text="Add new file">
                <IconButton ref={focus} onClick={() => setIsOpen(true)} icon={PlusIcon} aria-label='Add file' />
            </Tooltip>
            <Dialog
                returnFocusRef={focus}
                isOpen={isOpen}
                onDismiss={() => setIsOpen(false)}
            >
                <Dialog.Header>
                    Create new query snippet
                </Dialog.Header>
                <Box p={3} width={"100%"}>
                    <FormControl required>
                        <FormControl.Label>Name</FormControl.Label>
                        <TextInput block />
                    </FormControl>
                    <FormControl>
                        <FormControl.Label>Description</FormControl.Label>
                        <TextInput block />
                    </FormControl>
                    <Button block variant='primary' sx={{ my: 2 }}>Create</Button>
                </Box>
            </Dialog>
        </Box>
    )
}