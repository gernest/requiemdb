import { Box, NavList } from "@primer/react"
import { DotIcon, HomeIcon, CodeIcon, GearIcon, Icon } from "@primer/octicons-react";

export const Sidebar = () => {
    return (
        <Box>

            <NavList>
                <NavList>
                    <NavList.Item href="/" aria-current="page">
                        <NavList.LeadingVisual>
                            <HomeIcon />
                        </NavList.LeadingVisual>
                        Home
                    </NavList.Item>
                    <NavList.Item href="/about">
                        <NavList.LeadingVisual>
                            <DotIcon />
                        </NavList.LeadingVisual>
                        Console
                    </NavList.Item>
                    <NavList.Item href="/about">
                        <NavList.LeadingVisual>
                            <CodeIcon />
                        </NavList.LeadingVisual>
                        Snippets
                    </NavList.Item>
                    <NavList.Item href="/contact">
                        <NavList.LeadingVisual>
                            <GearIcon />
                        </NavList.LeadingVisual>
                        Settings
                    </NavList.Item>
                </NavList>
            </NavList>
        </Box>
    )
}

