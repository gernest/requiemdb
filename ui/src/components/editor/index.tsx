import * as monaco from 'monaco-editor';
import { useEffect, useRef, useState } from 'react';
import { Box, BoxProps } from '@primer/react';

const welcome = `import { metrics, render } from "@requiemdb/rq";

/**
 *  Instant Vectors
 */
render(
    metrics()
        .name("http_requests_total")
        .query(),
)
`

export const Editor = (props: BoxProps) => {
    const [editor, setEditor] = useState<monaco.editor.IStandaloneCodeEditor | null>(null);
    const hostRef = useRef<HTMLDivElement>(null);
    useEffect(() => {
        if (hostRef) {
            setEditor((editor) => {
                if (editor) return editor;
                return monaco.editor.create(hostRef.current!, {
                    language: "typescript",
                    fontSize: 18,
                    value: welcome,
                    automaticLayout: true,
                    minimap: { enabled: false }
                });
            });
        }
        return () => editor?.dispose();
    }, [hostRef.current]);

    return (
        <Box ref={hostRef} overflow={"hidden"}{...props} />
    )
}