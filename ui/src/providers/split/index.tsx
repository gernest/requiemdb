import {
    createContext,
    PropsWithChildren,
    useContext,
    useRef, useState, Ref, useEffect, MouseEvent as ReactEvent
} from "react"


type ContextProps = {
    clientHeight: number
    setClientHeight: (height: number) => void
    clientWidth: number
    setClientWidth: (width: number) => void
    xDividerPos: Ref<number>
    yDividerPos: Ref<number>
    onMouseHoldDown: (e: ReactEvent) => void
}


type Props = {}

export const SplitContext = createContext<ContextProps | null>(null);


export const SplitProvider = ({ children }: PropsWithChildren<Props>) => {
    const [clientHeight, setClientHeight] = useState<number>(0);
    const [clientWidth, setClientWidth] = useState<number>(0);
    const yDividerPos = useRef<number | null>(0);
    const xDividerPos = useRef<number | null>(0);
    const onMouseHoldDown = (e: ReactEvent) => {
        yDividerPos.current = e.clientY;
        xDividerPos.current = e.clientX;
    };

    const onMouseHoldUp = () => {
        yDividerPos.current = null;
        xDividerPos.current = null;
    };
    const onMouseHoldMove = (e: MouseEvent) => {
        if (!yDividerPos.current && !xDividerPos.current) {
            return;
        }

        setClientHeight(clientHeight + e.clientY - yDividerPos.current!);
        setClientWidth(clientWidth + e.clientX - xDividerPos.current!);
        yDividerPos.current = e.clientY;
        xDividerPos.current = e.clientX;
    };

    useEffect(() => {
        document.addEventListener("mouseup", onMouseHoldUp);
        document.addEventListener("mousemove", onMouseHoldMove);

        return () => {
            document.removeEventListener("mouseup", onMouseHoldUp);
            document.removeEventListener("mousemove", onMouseHoldMove);
        };
    });

    return (
        <SplitContext.Provider value={{
            clientHeight, setClientHeight,
            clientWidth, setClientWidth,
            yDividerPos, xDividerPos,
            onMouseHoldDown
        }}>
            {children}
        </SplitContext.Provider>
    )
}


export const useSplit = () => {
    return useContext(SplitContext)!
}