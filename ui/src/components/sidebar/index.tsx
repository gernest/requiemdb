import { Box, Octicon, Tooltip, TooltipProps } from "@primer/react"
import { DotIcon, HomeIcon, CodeIcon, GearIcon, Icon } from "@primer/octicons-react";

export const Sidebar = () => {
    return (
        <Box height={"100%"} padding={1}
            backgroundColor={'header.bg'}
            color={'header.text'}
            display={"grid"}
            alignItems={"center"}
            gridTemplateRows={"auto auto auto  1fr auto"}
        >

            <Item icon={HomeIcon} text="Dashboard" />
            <Item icon={DotIcon} text="Console" />
            <Item icon={CodeIcon} text="Snippets" />
            <Box></Box>
            <Item icon={GearIcon} text="Settings" />
        </Box>
    )
}

type Props = {
    icon: Icon,
} & TooltipProps

const Item = ({ icon, ...rest }: Props) => {
    return (
        <Box
            display={"flex"} alignItems={"center"} flexDirection={"column"} padding={2}
        >
            <Tooltip {...rest} direction="e">
                <Octicon icon={icon} size={"medium"} />
            </Tooltip>
        </Box>
    )
}