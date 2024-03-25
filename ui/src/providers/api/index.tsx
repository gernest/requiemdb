import { createContext, PropsWithChildren, useContext, useState } from "react";
import { createRQClient, RQClient } from "../../rq";

type ContextProps = {
    client: RQClient
}

type Props = {}

export const ApiContext = createContext<ContextProps>({ client: createRQClient() })

export const ApiProvider = ({ children }: PropsWithChildren<Props>) => {
    const [client, setClient] = useState<RQClient>(createRQClient());
    return (
        <ApiContext.Provider value={{ client }}>
            {children}
        </ApiContext.Provider>
    )
}

export const useApi = () => {
    return useContext(ApiContext)!
}