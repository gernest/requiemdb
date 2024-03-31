import { Box, NavList } from "@primer/react"
import { DotIcon, GearIcon } from "@primer/octicons-react";
import { NavItem } from "../nav";

export const Sidebar = () => {
    return (
        <Box>
            <NavList>
                <NavList>

                    <NavItem to={"/console"}>
                        <NavList.LeadingVisual>
                            <DotIcon />
                        </NavList.LeadingVisual>
                        Console
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

