import { useMatch, useResolvedPath, useNavigate, To } from 'react-router-dom'
import { TreeView, TreeViewItemProps } from '@primer/react'

export const TreeLinkItem = ({ id, to, children }: TreeViewItemProps & { to: To }) => {
    const navigate = useNavigate()
    const resolved = useResolvedPath(to)
    const isCurrent = useMatch({ path: resolved.pathname, end: true })
    return (
        <TreeView.Item
            id={id}
            current={isCurrent != null}
            onSelect={() => navigate(to)}
        >
            {children}
        </TreeView.Item>
    )
}

