import { Box } from "@primer/react"
import { TitleBar } from "./title"


export const Sidebar = () => {
    return (
        <Box>
            <aside>
                <Box display={"flex"} flexDirection={"column"} width={"100%"} flexGrow={1}>
                    <TitleBar />
                </Box>
            </aside>
        </Box>
    )
}