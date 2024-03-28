import { Box, NavList } from "@primer/react"
import { DotIcon, HomeIcon, CodeIcon, GearIcon } from "@primer/octicons-react";
import { NavItem } from "../nav";

export const Sidebar = () => {
    return (
        <Box>
            <NavList>
                <NavList>
                    <NavItem to={"/dashboard"}>
                        <NavList.LeadingVisual>
                            <HomeIcon />
                        </NavList.LeadingVisual>
                        Dashboard
                    </NavItem>
                    <NavItem to={"/console"}>
                        <NavList.LeadingVisual>
                            <DotIcon />
                        </NavList.LeadingVisual>
                        Console
                    </NavItem>
                    <NavItem to={"/snippets"}>
                        <NavList.LeadingVisual>
                            <CodeIcon />
                        </NavList.LeadingVisual>
                        Snippets
                    </NavItem>
                    <NavItem to={"/settings"}>
                        <NavList.LeadingVisual>
                            <GearIcon />
                        </NavList.LeadingVisual>
                        Settings
                    </NavItem>
                </NavList>
            </NavList>
        </Box>
    )
}

