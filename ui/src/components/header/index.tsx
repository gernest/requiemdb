import { Box, Header, Text } from "@primer/react"
import { Requiem } from "../rq";


export const MainHeader = () => {
    return (
        <Box
            sx={{
                position: 'sticky',
                top: 0,
                height: 64,
                backgroundColor: 'canvas.subtle',
                borderBottom: '1px solid',
                borderColor: 'border.default',
                zIndex: 1,
            }}
        >
            <Header sx={{ color: "fg.default", backgroundColor: 'canvas.subtle' }} >
                <Header.Item>
                    <Box>
                        <Requiem width={32} />
                    </Box>
                    <Text sx={{ fontWeight: "semibold" }}>RequiemDB</Text>
                </Header.Item>
            </Header>
        </Box>
    )
}