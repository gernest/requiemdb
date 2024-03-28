import { Link, useMatch, useResolvedPath, To } from 'react-router-dom'
import { NavList, NavListItemProps } from '@primer/react'

export const NavItem = ({ to, children }: NavListItemProps & { to: To }) => {
    const resolved = useResolvedPath(to)
    const isCurrent = useMatch({ path: resolved.pathname, end: true })
    return (
        <NavList.Item as={Link} to={to} aria-current={isCurrent ? 'page' : undefined}>
            {children}
        </NavList.Item>
    )
}