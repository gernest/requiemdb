import { NavList } from "@primer/react"
import { DotIcon, HomeIcon, CodeIcon, GearIcon } from "@primer/octicons-react";

export const Sidebar = () => {
    return (
        <NavList>
            <NavList.Item>
                <NavList.LeadingVisual>
                    <HomeIcon />
                </NavList.LeadingVisual>
                Home
            </NavList.Item>
            <NavList.Item>
                <NavList.LeadingVisual>
                    <DotIcon />
                </NavList.LeadingVisual>
                Console
            </NavList.Item>
            <NavList.Item>
                <NavList.LeadingVisual>
                    <CodeIcon />
                </NavList.LeadingVisual>
                Snippets
            </NavList.Item>
            <NavList.Item>
                <NavList.LeadingVisual>
                    <GearIcon />
                </NavList.LeadingVisual>
                Settings
            </NavList.Item>
        </NavList>
    )
}