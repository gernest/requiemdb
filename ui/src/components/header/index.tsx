import { Box, Header, Text } from "@primer/react"
import { Requiem } from "../rq";


export const MainHeader = () => {
    return (
        <Header sx={{ color: "fg.default", backgroundColor: 'canvas.subtle' }} >
            <Header.Item>
                <Box>
                    <Requiem width={32} />
                </Box>
                <Text sx={{ fontWeight: "semibold" }}>RequiemDB</Text>
            </Header.Item>
        </Header>
    )
}