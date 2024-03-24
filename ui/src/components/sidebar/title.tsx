import { Box } from "@primer/react"
import { Requiem } from ".."

export const TitleBar = () => {
    return (
        <Box px={2} py={2}>
            <Box display={"flex"} alignItems={"center"}>
                <Box display={"flex"} alignItems={"center"}>
                    <Requiem width={30} />
                </Box>
                <Box display={"flex"} alignItems={"center"}
                    paddingLeft={2} fontSize={1} fontWeight={"semibold"}
                    position={"relative"} top={-1}>
                    RequiemDB
                </Box>
            </Box>
        </Box>
    )
}